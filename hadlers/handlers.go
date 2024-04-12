package hadlers

import (
	"avito_test_task/cache"
	"avito_test_task/db"
	"avito_test_task/models"
	"avito_test_task/service"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type BannerHandler struct {
	bannerService    *service.BannerService
	bannerTagService *service.BannerTagService
}

func NewBannerHandler(bannerService *service.BannerService, bannerTagService *service.BannerTagService) *BannerHandler {
	return &BannerHandler{
		bannerService:    bannerService,
		bannerTagService: bannerTagService,
	}
}

func (b *BannerHandler) Create(c *gin.Context) {

	token := c.GetHeader("token")
	if token == "" {
		c.Status(http.StatusUnauthorized)
		return
	} else if token != "admin_token" {
		c.Status(http.StatusForbidden)
		return
	}

	var request models.BannerRequestBody
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bannerExists, err := b.bannerService.IsBannerWithFeatureAndTagExists(request.FeatureID, request.TagIds, 0)
	if bannerExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "banner with this tag_id and feature_id already exists"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	banner, err := b.bannerService.CreateBannerWithTags(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"banner_id": banner.BannerID})

}

func (b *BannerHandler) Delete(c *gin.Context) {
	token := c.GetHeader("token")
	if token == "" {
		c.Status(http.StatusUnauthorized)
		return
	} else if token != "admin_token" {
		c.Status(http.StatusForbidden)
		return
	}
	idString := c.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	result := b.bannerService.DeleteByID(id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	c.Status(http.StatusNoContent)

}

// здесь должна быть транзакция, тк если указаны неверные теги, то баннер обновляется, теги удаляются и новые не добавляются
func (b *BannerHandler) Update(c *gin.Context) {
	token := c.GetHeader("token")
	if token == "" {
		c.Status(http.StatusUnauthorized)
		return
	} else if token != "admin_token" {
		c.Status(http.StatusForbidden)
		return
	}
	idString := c.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var request models.BannerRequestBody
	if err = c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	oldBanner, err := b.bannerService.FindByID(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Status(http.StatusNotFound)
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	bannerExists, err := b.bannerService.IsBannerWithFeatureAndTagExists(request.FeatureID, request.TagIds, oldBanner.BannerID)
	if bannerExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "banner with this tag_id and feature_id already exists"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	tx := db.DB.Begin()
	resultUpdate := b.bannerService.UpdateBanner(oldBanner, &request)

	if resultUpdate.Error != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": resultUpdate.Error.Error()})
		return
	}

	resultDelete := b.bannerTagService.DeleteByBannerID(oldBanner.BannerID)

	if resultDelete.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": resultDelete.Error.Error()})
		return
	}
	if resultDelete.RowsAffected == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	for _, tagID := range request.TagIds {
		if err = b.bannerTagService.Create(&models.BannerTag{
			BannerID: id,
			TagID:    tagID,
		}); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.Status(http.StatusOK)
}

func (b *BannerHandler) Get(c *gin.Context) {

	token := c.GetHeader("token")
	if token == "" {
		c.Status(http.StatusUnauthorized)
		return
	} else if token != "admin_token" {
		c.Status(http.StatusForbidden)
		return
	}
	featureId, _ := strconv.ParseUint(c.Query("feature_id"), 10, 64)
	tagId, _ := strconv.ParseUint(c.Query("tag_id"), 10, 64)
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))

	cacheKey := fmt.Sprintf("banners %d %d %d %d", featureId, tagId, limit, offset)
	if cachedBanners, found := cache.Get(cacheKey); found {
		c.JSON(http.StatusOK, cachedBanners)
		return
	}

	banners, err := b.bannerService.GetBanners(featureId, tagId, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	cache.Set(cacheKey, banners)
	c.JSON(http.StatusOK, banners)

}

func (b *BannerHandler) GetUserBanners(c *gin.Context) {
	token := c.GetHeader("token")
	if token == "" {
		c.Status(http.StatusUnauthorized)
		return
	} else if token != "admin_token" && token != "user_token" {
		c.Status(http.StatusForbidden)
		return
	}

	tagIdString := c.Query("tag_id")
	tagId, err := strconv.ParseUint(tagIdString, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tag_id"})
		return
	}

	featureIdString := c.Query("feature_id")
	featureId, err := strconv.ParseUint(featureIdString, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid feature_id"})
		return
	}
	//TODO: check the when use_last_revision is not present value is false
	useLastRevision, _ := strconv.ParseBool(c.Query("use_last_revision"))
	cacheKey := fmt.Sprintf("bannerTag %d %d", featureId, tagId)
	if !useLastRevision {
		if cachedBannerTag, found := cache.Get(cacheKey); found {
			c.JSON(http.StatusOK, cachedBannerTag.(models.Banner).Content)
			return
		}
	}
	banner, err := b.bannerService.GetUserBanner(featureId, tagId, token)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Status(http.StatusNotFound)
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	cache.Set(cacheKey, banner)
	c.JSON(http.StatusOK, banner.Content)

}
