package routers

import (
	"avito_test_task/models"
	"avito_test_task/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoutes(r *gin.Engine) {
	BannerHandler := NewBannerHandler(service.NewBannerService(), service.NewBannerTagService())

	bannerGroup := r.Group("/banner")
	{
		bannerGroup.GET("", temp)
		bannerGroup.POST("", BannerHandler.CreateBanner)
		bannerGroup.PATCH("/:id", temp)
		bannerGroup.DELETE("/:id", temp)
	}
	r.GET("/user-banner", temp)
}

func temp(c *gin.Context) {
	c.JSON(http.StatusOK, models.Tag{
		TagID:       0,
		Description: "test",
	})
}
