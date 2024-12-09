package cmd

import (
	"fmt"
	"github.com/caovanhoang63/hiholive/services/video/composer"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx/components/ginc"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx/components/ginc/middlewares"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx/components/gormc"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx/components/jwtc"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"net/http"
	"os"
)

func newServiceCtx() srvctx.ServiceContext {
	return srvctx.NewServiceContext(
		srvctx.WithName("Demo Microservices"),
		srvctx.WithComponent(ginc.NewGin(core.KeyCompGIN)),
		srvctx.WithComponent(gormc.NewGormDB(core.KeyCompMySQL, "")),
		srvctx.WithComponent(jwtc.NewJWT(core.KeyCompJWT)),
		srvctx.WithComponent(core.NewConfig()),
	)
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start service",
	Run: func(cmd *cobra.Command, args []string) {
		serviceCtx := newServiceCtx()

		logger := srvctx.GlobalLogger().GetLogger("service")

		if err := serviceCtx.Load(); err != nil {
			logger.Fatal(err)
		}

		ginComp := serviceCtx.MustGet(core.KeyCompGIN).(core.GINComponent)

		router := ginComp.GetRouter()
		router.Use(gin.Recovery(), middlewares.Logger(serviceCtx), middlewares.Recovery(serviceCtx))

		router.Use(middlewares.Cors())
		router.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "pong"})
		})

		v1 := router.Group("/v1")

		SetupRoutes(v1, serviceCtx)

		if err := router.Run(fmt.Sprintf(":%d", ginComp.GetPort())); err != nil {
			logger.Fatal(err)
		}
	},
}

func SetupRoutes(router *gin.RouterGroup, serviceCtx srvctx.ServiceContext) {
	channelService := composer.ComposeChannelAPIService(serviceCtx)
	ac := composer.ComposeAuthRPCClient(serviceCtx)
	uc := composer.ComposeUserRPCClient(serviceCtx)
	channel := router.Group("/channel")

	channel.POST("", middlewares.RequireAuth(ac), middlewares.Authorize(uc, "viewer"), channelService.CreateChannel())
}

func Execute() {
	rootCmd.AddCommand(outEnvCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
