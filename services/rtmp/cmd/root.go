package cmd

import (
	"fmt"
	"github.com/caovanhoang63/hiholive/services/rtmp/components/rtmpc"
	"github.com/caovanhoang63/hiholive/services/rtmp/composer"
	"github.com/caovanhoang63/hiholive/shared/golang/core"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/ginc"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/jwtc"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/pubsub"
	rabbitpubsub "github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/pubsub/rabbitmq"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/redisc"
	"github.com/spf13/cobra"
	"github.com/yutopp/go-rtmp"
	"io"
	"net"
	"time"

	"os"
)

func newServiceCtx() srvctx.ServiceContext {
	return srvctx.NewServiceContext(
		srvctx.WithName("RTMP Service"),
		srvctx.WithComponent(ginc.NewGin(core.KeyCompGIN)),
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
		rd := serviceCtx.MustGet(core.KeyRedis).(core.RedisComponent)
		ps := serviceCtx.MustGet(core.KeyCompRabbitMQ).(pubsub.Pubsub)

		logger := srvctx.GlobalLogger().GetLogger("Rtmp Service")
		if err := serviceCtx.Load(); err != nil {
			logger.Fatal(err)
		}

		tcpAddr, err := net.ResolveTCPAddr("tcp", ":1935")
		if err != nil {
			logger.Panicf("Failed: %+v", err)
		}
		logger.Info("Start TCP port")
		listener, err := net.ListenTCP("tcp", tcpAddr)
		if err != nil {
			logger.Panicf("Failed: %+v", err)
		}

		// Setup TCP connection
		relayService := rtmpc.NewRelayService()
		srv := rtmp.NewServer(&rtmp.ServerConfig{
			OnConnect: func(conn net.Conn) (io.ReadWriteCloser, *rtmp.ConnConfig) {
				return conn, &rtmp.ConnConfig{
					Handler: rtmpc.NewHandler(relayService, rd.GetClient(), composer.ComposeHlsRPCClient(serviceCtx), ps),
					Timeout: time.Second * 5,
					ControlState: rtmp.StreamControlStateConfig{
						DefaultBandwidthWindowSize: 6 * 1024 * 1024 / 8,
					},
					Logger: nil,
				}
			},
		})

		if err = srv.Serve(listener); err != nil {
			logger.Error(err)
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
