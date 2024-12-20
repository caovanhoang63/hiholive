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
	ctgService := composer.ComposeCategoryApiService(serviceCtx)

	ac := composer.ComposeAuthRPCClient(serviceCtx)
	uc := composer.ComposeUserRPCClient(serviceCtx)

	channel := v1.Group("/channel")
	channel.POST("", middlewares.RequireAuth(ac), middlewares.Authorize(uc, "viewer"), channelService.CreateChannel())

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
