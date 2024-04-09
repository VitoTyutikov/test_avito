package routers

import (
	"avito_test_task/models"
	"avito_test_task/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoutes(r *gin.Engine) {
	bannerHandler := NewBannerHandler(service.NewBannerService(), service.NewBannerTagService())

	bannerGroup := r.Group("/banner")
	{
		bannerGroup.GET("", temp)
		bannerGroup.POST("", bannerHandler.Create)
		bannerGroup.PATCH("/:id", temp)
		bannerGroup.DELETE("/:id", bannerHandler.Delete)
	}
	r.GET("/user-banner", temp)
}

func temp(c *gin.Context) {
	c.JSON(http.StatusOK, models.Tag{
		TagID:       0,
		Description: "test",
	})
}
