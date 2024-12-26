package cmd

import (
	"github.com/caovanhoang63/hiholive/services/video/strmcomposer"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx/components/ginc/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, serviceCtx srvctx.ServiceContext) {
	v1 := router.Group("v1")
	channelService := strmcomposer.ComposeChannelAPIService(serviceCtx)
	streamService := strmcomposer.ComposeStreamAPIService(serviceCtx)
	settingService := strmcomposer.ComposeSystemSettingApiService(serviceCtx)
	ctgService := strmcomposer.ComposeCategoryApiService(serviceCtx)

	ac := strmcomposer.ComposeAuthRPCClient(serviceCtx)
	uc := strmcomposer.ComposeUserRPCClient(serviceCtx)

	channelPrv := v1.Group("/channel")
	channelPrv.POST("", middlewares.RequireAuth(ac), middlewares.Authorize(uc, "viewer"), channelService.CreateChannel())

	channelPub := v1.Group("/channel")
	channelPub.GET(":id", channelService.FindChannelById())
	channelPub.GET("", channelService.FindChannels())

	userPub := v1.Group("/user")
	userPub.GET("/:id/channel", channelService.FindChannelById())

	stream := v1.Group("/stream")
	stream.POST("", middlewares.RequireAuth(ac), middlewares.Authorize(uc, "streamer"), streamService.CreateStream())

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
}
