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

	ac := composer.ComposeAuthRPCClient(serviceCtx)
	uc := composer.ComposeUserRPCClient(serviceCtx)

	channel := v1.Group("/channel")
	channel.POST("", middlewares.RequireAuth(ac), middlewares.Authorize(uc, "viewer"), channelService.CreateChannel())

	stream := v1.Group("/stream")
	stream.POST("", middlewares.RequireAuth(ac), middlewares.Authorize(uc, "streamer"), streamService.CreateStream())
}
