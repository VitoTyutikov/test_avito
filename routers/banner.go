package routers

import (
	"avito_test_task/models"
	"avito_test_task/service"
	"github.com/gin-gonic/gin"
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
	banner, err := b.bannerService.Create(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, tagID := range request.TagIds {
		if err := b.bannerTagService.Create(&models.BannerTag{
			BannerID: banner.BannerID,
			TagID:    tagID,
		}); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
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

//func CheckToken(token, requiredToken string) bool {
//	if token == "" {
//		return false
//	}
//	if requiredToken == "admin_token"{
//		return token == "admin_token"
//	} else if requiredToken == "user_token"{
//		return token == "user_token" || token == "admin_token"
//
//	}
//}
