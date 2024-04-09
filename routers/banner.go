package routers

import (
	"avito_test_task/models"
	"avito_test_task/service"
	"github.com/gin-gonic/gin"
	"net/http"
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

func (b *BannerHandler) CreateBanner(c *gin.Context) {
	token := c.GetHeader("token")
	if token == "" {
		c.Status(http.StatusUnauthorized)
		return
	}
	if token != "admin_token" {
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
	//c.Status(http.StatusCreated)

}
