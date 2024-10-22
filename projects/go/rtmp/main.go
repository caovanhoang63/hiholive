package main

import (
	"context"
	"github.com/yutopp/go-rtmp"
	"go.uber.org/zap"
	"hiholive/projects/go/rtmp/component/appContext"
	"hiholive/shared/go/srvctx"
	"io"
	"net"
)

func main() {
	// Setup dependencies
	l := srvctx.NewZapLogger(context.Background(), zap.DebugLevel)
	appCtx := appContext.NewAppContextRTPM(l)
	log := appCtx.GetLogger()
	// Setup dependencies

	// Setup TCP connection
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1935")
	if err != nil {
		log.FatalWithFields("Failed: %+v", srvctx.Field{"Error": err})
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.FatalWithFields("Failed: %+v", srvctx.Field{"Error": err})
	}
	// Setup TCP connection

	relayService := NewRelayService()

	srv := rtmp.NewServer(&rtmp.ServerConfig{
		OnConnect: func(conn net.Conn) (io.ReadWriteCloser, *rtmp.ConnConfig) {
			h := &Handler{
				relayService: relayService,
			}

			return conn, &rtmp.ConnConfig{
				Handler: h,

				ControlState: rtmp.StreamControlStateConfig{
					DefaultBandwidthWindowSize: 6 * 1024 * 1024 / 8,
				},
			}
		},
	})

	if err = srv.Serve(listener); err != nil {
		log.FatalWithFields("Failed: %+v", srvctx.Field{"Error": err})
	}
}
