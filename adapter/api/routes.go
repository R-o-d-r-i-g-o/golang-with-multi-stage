package api

import (
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func routes(router *gin.Engine) *gin.Engine {
	video := videoInfo()

	api := router.Group("api/video")
	{
		api.POST("/", video.CreateVideo)
		api.GET("/", video.GetVideoInfo)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
