package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"hiholive/projects/go/user/composer"
	"hiholive/shared/go/shared"
	"hiholive/shared/go/srvctx"
	"hiholive/shared/go/srvctx/components/ginc"
	"hiholive/shared/go/srvctx/components/ginc/middlewares"
	"hiholive/shared/go/srvctx/components/gormc"
	"hiholive/shared/go/srvctx/components/jwtc"
	"net/http"
	"os"
)

func newServiceCtx() srvctx.ServiceContext {
	return srvctx.NewServiceContext(
		srvctx.WithName("Demo Microservices"),
		srvctx.WithComponent(ginc.NewGin(shared.KeyCompGIN)),
		srvctx.WithComponent(gormc.NewGormDB(shared.KeyCompMySQL, "")),
		srvctx.WithComponent(jwtc.NewJWT(shared.KeyCompJWT)),
		srvctx.WithComponent(NewConfig()),
	)
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start service",
	Run: func(cmd *cobra.Command, args []string) {
		serviceCtx := newServiceCtx()

		logger := srvctx.GlobalLogger().GetLogger("service")

		// Make some delay for DB ready (migration)
		// remove it if you already had your own DB

		if err := serviceCtx.Load(); err != nil {
			logger.Fatal(err.Error())
		}

		ginComp := serviceCtx.MustGet(shared.KeyCompGIN).(shared.GINComponent)

		router := ginComp.GetRouter()
		router.Use(gin.Recovery(), middlewares.Logger(serviceCtx), middlewares.Recovery(serviceCtx))

		router.Use(middlewares.Cors())
		router.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "pong"})
		})

		v1 := router.Group("/v1")

		SetupRoutes(v1, serviceCtx)

		if err := router.Run(fmt.Sprintf(":%d", ginComp.GetPort())); err != nil {
			logger.Fatal(err.Error())
		}
	},
}

func SetupRoutes(router *gin.RouterGroup, serviceCtx srvctx.ServiceContext) {
	userService := composer.ComposeUserAPIService(serviceCtx)

	tasks := router.Group("/user")
	{
		tasks.GET("", userService.GetUserProfile())
	}
}

func Execute() {
	rootCmd.AddCommand(outEnvCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
