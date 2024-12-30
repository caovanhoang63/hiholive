package cmd

import (
	"github.com/caovanhoang63/hiholive/services/video/videocomposer"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/ginc/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, serviceCtx srvctx.ServiceContext) {
	v1 := router.Group("v1")
	ac := videocomposer.ComposeAuthRPCClient(serviceCtx)
	uc := videocomposer.ComposeUserRPCClient(serviceCtx)

	channelService := videocomposer.ComposeChannelAPIService(serviceCtx, uc)
	streamService := videocomposer.ComposeStreamAPIService(serviceCtx)
	settingService := videocomposer.ComposeSystemSettingApiService(serviceCtx)
	ctgService := videocomposer.ComposeCategoryApiService(serviceCtx)
	uploadService := videocomposer.ComposeUploadAPIService(serviceCtx)

	channelPrv := v1.Group("/channel")
	channelPrv.POST("", middlewares.RequireAuth(ac), middlewares.Authorize(uc, "viewer"), channelService.CreateChannel())

	channelPub := v1.Group("/channel")
	channelPub.GET(":id", channelService.FindChannelById())
	channelPub.GET("user/:username", channelService.FindChannelByUserName())
	channelPub.GET("", channelService.FindChannels())

	userPub := v1.Group("/user")
	userPub.GET("/:id/channel", channelService.FindChannelById())

	streamPrv := v1.Group("/stream")
	streamPrv.POST("", middlewares.RequireAuth(ac), middlewares.Authorize(uc, "streamer"), streamService.CreateStream())

	streamPub := v1.Group("/stream")
	streamPub.GET(":id", streamService.GetStreamById())
	streamPub.GET("", streamService.FindStreams())

	settingPrv := v1.Group("/setting")
	settingPrv.Use(middlewares.RequireAuth(ac), middlewares.Authorize(uc, "admin"))
	settingPrv.POST("", settingService.CreateSystemSetting())
	settingPrv.PATCH("/:name", settingService.UpdateSystemSetting())

	settingPublic := v1.Group("/setting")
	settingPublic.GET("", settingService.FindSystemSetting())
	settingPublic.GET("/:name", settingService.FindSystemSettingByName())

	ctgPrv := v1.Group("/category")
	ctgPrv.Use(middlewares.RequireAuth(ac), middlewares.Authorize(uc, "admin"))
	ctgPrv.POST("", ctgService.CreateCategory())
	ctgPrv.PATCH("/:id", ctgService.UpdateCategory())
	ctgPrv.DELETE("/:id", ctgService.DeleteCategory())

	ctgPublic := v1.Group("/category")
	ctgPublic.GET("", ctgService.FindCategories())
	ctgPublic.GET("/:id", ctgService.FindCategoryById())

	uploadPub := v1.Group("/upload/image")
	uploadPub.POST("", uploadService.UploadImage())
}
