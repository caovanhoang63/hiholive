package cmd

import (
	"fmt"
	"github.com/caovanhoang63/hiholive/services/hls_mux/component/ffmpegc"
	"github.com/caovanhoang63/hiholive/shared/golang/core"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/ginc"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/ginc/middlewares"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/gormc"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/jwtc"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/pubsub/rabbitmq"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/redisc"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	"os"
)

func newServiceCtx() srvctx.ServiceContext {
	return srvctx.NewServiceContext(
		srvctx.WithName("Demo Microservices"),
		srvctx.WithComponent(ginc.NewGin(core.KeyCompGIN)),
		srvctx.WithComponent(gormc.NewGormDB(core.KeyCompMySQL, "")),
		srvctx.WithComponent(jwtc.NewJWT(core.KeyCompJWT)),
		srvctx.WithComponent(core.NewConfig()),
		srvctx.WithComponent(redisc.NewRedis(core.KeyRedis)),
		srvctx.WithComponent(rabbitpubsub.NewRabbitPubSub(core.KeyCompRabbitMQ)),
	)
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start service",
	Run: func(cmd *cobra.Command, args []string) {
		serviceCtx := newServiceCtx()

		logger := srvctx.GlobalLogger().GetLogger("Hls service")
		if err := serviceCtx.Load(); err != nil {
			logger.Fatal(err)
		}

		ginComp := serviceCtx.MustGet(core.KeyCompGIN).(core.GINComponent)
		rd := serviceCtx.MustGet(core.KeyRedis).(core.RedisComponent)
		router := ginComp.GetRouter()

		router.Static("/static", "hls_output")
		router.Use(middlewares.Cors())

		router.GET("ping", func(c *gin.Context) {
			c.JSON(200, "pong")
		})

		ffmpeg := ffmpegc.NewFfmpeg(serviceCtx).WithConfig(ffmpegc.NewFfmpegConfig("./hls_output", rd.GetClient()))

		go func() {
			defer core.AppRecover()
			StartGRPCServices(ffmpeg, serviceCtx)
		}()

		go func() {
			defer core.AppRecover()
			StartSubscriberServices(ffmpeg, serviceCtx)
		}()
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
