package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yutopp/go-rtmp"
	"hiholive/projects/go/rtmp/components/rtmpc"
	"hiholive/shared/go/shared"
	"hiholive/shared/go/srvctx"
	"hiholive/shared/go/srvctx/components/ginc"
	"hiholive/shared/go/srvctx/components/gormc"
	"hiholive/shared/go/srvctx/components/jwtc"
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
		logger.Info("Start TCP port")
		listener, err := net.ListenTCP("tcp", tcpAddr)
		if err != nil {
			logger.Error(err)
		}
		// Setup TCP connection
		relayService := rtmpc.NewRelayService()
		rtmpHandler := rtmpc.NewHandler(relayService)
		srv := rtmp.NewServer(&rtmp.ServerConfig{
			OnConnect: func(conn net.Conn) (io.ReadWriteCloser, *rtmp.ConnConfig) {
				h := rtmpHandler
				return conn, &rtmp.ConnConfig{
					Handler: h,
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
