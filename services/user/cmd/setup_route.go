package cmd

import (
	"github.com/caovanhoang63/hiholive/services/user/usercomposer"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/ginc/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, serviceCtx srvctx.ServiceContext) {
	v1 := router.Group("v1")
	userService := usercomposer.ComposeUserAPIService(serviceCtx)

	userPub := v1.Group("user")
	userPub.GET(":id", userService.GetUserById())
	userPub.GET("", userService.ListUser())

	userPrv := v1.Group("user")
	userPrv.Use(middlewares.RequireAuth(usercomposer.ComposeAuthRPCClient(serviceCtx)))
	userPrv.GET("profile", userService.GetUserProfile())
}
