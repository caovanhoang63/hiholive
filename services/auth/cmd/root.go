package cmd

import (
	"fmt"
	"github.com/caovanhoang63/hiholive/shared/golang/core"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/ginc"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/ginc/middlewares"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/gormc"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/jwtc"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/pubsub/rabbitmq"
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
		srvctx.WithComponent(rabbitpubsub.NewRabbitPubSub(core.KeyCompRabbitMQ)),
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

		go StartGRPCServices(serviceCtx)

		SetupRoutes(router.Group(""), serviceCtx)

		if err := router.Run(fmt.Sprintf(":%d", ginComp.GetPort())); err != nil {
			logger.Fatal(err)
		}
	},
}

func Execute() {
	rootCmd.AddCommand(outEnvCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
