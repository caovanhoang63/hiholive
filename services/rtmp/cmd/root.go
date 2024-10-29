package cmd

import (
	"fmt"
	"github.com/caovanhoang63/hiholive/services/rtmp/components/rtmpc"
	"github.com/caovanhoang63/hiholive/shared/go/shared"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx/components/ginc"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx/components/gormc"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx/components/jwtc"
	"github.com/spf13/cobra"
	"github.com/yutopp/go-rtmp"

	"io"
	"net"
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
		logger := srvctx.GlobalLogger().GetLogger("Rtmp Service")
		// Setup TCP connection
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
					Handler: rtmpc.NewHandler(relayService),
					ControlState: rtmp.StreamControlStateConfig{
						DefaultBandwidthWindowSize: 6 * 1024 * 1024 / 8,
					},
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
