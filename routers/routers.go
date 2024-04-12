package routers

import (
	"avito_test_task/hadlers"
	"avito_test_task/service"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	bannerHandler := hadlers.NewBannerHandler(service.NewBannerService(), service.NewBannerTagService())

	bannerGroup := r.Group("/banner")
	{
		bannerGroup.GET("", bannerHandler.Get)
		bannerGroup.POST("", bannerHandler.Create)
		bannerGroup.PATCH("/:id", bannerHandler.Update)
		bannerGroup.DELETE("/:id", bannerHandler.Delete)
	}
	r.GET("/user_banner", bannerHandler.GetUserBanners)
}
