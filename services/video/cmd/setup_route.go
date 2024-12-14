package cmd

import (
	"github.com/caovanhoang63/hiholive/services/video/composer"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx/components/ginc/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, serviceCtx srvctx.ServiceContext) {
	v1 := router.Group("v1")
	channelService := composer.ComposeChannelAPIService(serviceCtx)
	streamService := composer.ComposeStreamAPIService(serviceCtx)
	settingService := composer.ComposeSystemSettingApiService(serviceCtx)

	ac := composer.ComposeAuthRPCClient(serviceCtx)
	uc := composer.ComposeUserRPCClient(serviceCtx)

	channel := v1.Group("/channel")
	channel.POST("", middlewares.RequireAuth(ac), middlewares.Authorize(uc, "viewer"), channelService.CreateChannel())

	stream := v1.Group("/stream")
	stream.POST("", middlewares.RequireAuth(ac), middlewares.Authorize(uc, "streamer"), streamService.CreateStream())

	settingPrv := v1.Group("/setting")
	settingPrv.Use(middlewares.RequireAuth(ac), middlewares.Authorize(uc, "admin"))
	settingPrv.POST("", settingService.CreateSystemSetting())
	settingPrv.PATCH("", settingService.UpdateSystemSetting())

	settingPublic := v1.Group("/setting")
	settingPublic.GET("", settingService.FindSystemSetting())
	settingPublic.GET("/:name", settingService.FindSystemSettingByName())
}
