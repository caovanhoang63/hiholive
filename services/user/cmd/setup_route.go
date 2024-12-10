package cmd

import (
	"github.com/caovanhoang63/hiholive/services/user/composer"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx/components/ginc/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, serviceCtx srvctx.ServiceContext) {
	v1 := router.Group("v1")
	userService := composer.ComposeUserAPIService(serviceCtx)

	tasks := v1.Group("user")
	{
		tasks.Use(middlewares.RequireAuth(composer.ComposeAuthRPCClient(serviceCtx)))
		tasks.GET(":id", userService.GetUserProfile())
		tasks.GET("", userService.ListUser())
	}
}
