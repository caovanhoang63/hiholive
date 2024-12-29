package cmd

import (
	"github.com/caovanhoang63/hiholive/services/auth/composer"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.RouterGroup, serviceCtx srvctx.ServiceContext) {
	v1 := router.Group("v1")
	authService := composer.ComposeAuthAPIService(serviceCtx)

	auth := v1.Group("auth")
	{
		auth.POST("register", authService.Register())
		auth.POST("login", authService.Login())
	}
}
