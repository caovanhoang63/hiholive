package cmd

import (
	"fmt"
	"github.com/caovanhoang63/hiholive/services/auth/common"
	"github.com/caovanhoang63/hiholive/services/hls_mux/component/ffmpegc"
	"github.com/caovanhoang63/hiholive/shared/go/shared"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx/components/ginc"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx/components/jwtc"
	"github.com/spf13/cobra"

	"os"
)

func newServiceCtx() srvctx.ServiceContext {
	return srvctx.NewServiceContext(
		srvctx.WithName("Demo Microservices"),
		srvctx.WithComponent(ginc.NewGin(shared.KeyCompGIN)),
		//srvctx.WithComponent(gormc.NewGormDB(shared.KeyCompMySQL, "")),
		srvctx.WithComponent(jwtc.NewJWT(shared.KeyCompJWT)),
		srvctx.WithComponent(NewConfig()),
		//srvctx.WithComponent(rabbitpubsub.NewRabbitPubSub(shared.KeyCompRabbitMQ)),
	)
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start service",
	Run: func(cmd *cobra.Command, args []string) {
		serviceCtx := newServiceCtx()

		logger := srvctx.GlobalLogger().GetLogger("Rtmp Service")
		if err := serviceCtx.Load(); err != nil {
			logger.Fatal(err)
		}

		//_ = serviceCtx.MustGet(shared.KeyCompRabbitMQ).(pubsub.Pubsub)
		ginComp := serviceCtx.MustGet(shared.KeyCompGIN).(common.GINComponent)
		router := ginComp.GetRouter()

		router.Static("/static", "hls_output")

		ffmpeg := ffmpegc.NewFfmpeg(serviceCtx).WithConfig(nil)

		go func() {
			ffmpeg.NewStream("test")
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
