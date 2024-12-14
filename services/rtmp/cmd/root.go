package cmd

import (
	"fmt"
	"github.com/caovanhoang63/hiholive/services/rtmp/components/rtmpc"
	"github.com/caovanhoang63/hiholive/services/rtmp/composer"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx/components/ginc"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx/components/jwtc"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx/components/redisc"
	"github.com/spf13/cobra"
	"github.com/yutopp/go-rtmp"
	"io"
	"net"
	"os"
)

func newServiceCtx() srvctx.ServiceContext {
	return srvctx.NewServiceContext(
		srvctx.WithName("Demo Microservices"),
		srvctx.WithComponent(ginc.NewGin(core.KeyCompGIN)),
		srvctx.WithComponent(jwtc.NewJWT(core.KeyCompJWT)),
		srvctx.WithComponent(core.NewConfig()),
		srvctx.WithComponent(redisc.NewRedis(core.KeyRedis)),
	)
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start service",
	Run: func(cmd *cobra.Command, args []string) {
		serviceCtx := newServiceCtx()
		rd := serviceCtx.MustGet(core.KeyRedis).(core.RedisComponent)

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
					Handler: rtmpc.NewHandler(relayService, rd.GetClient(), composer.ComposeHlsRPCClient(serviceCtx)),
					ControlState: rtmp.StreamControlStateConfig{
						DefaultBandwidthWindowSize: 6 * 1024 * 1024 / 8,
					},
					//Logger: log.StandardLogger(),
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
