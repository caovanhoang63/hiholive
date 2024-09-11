package main

import (
	"context"
	"github.com/yutopp/go-rtmp"
	"go.uber.org/zap"
	"hiholive/projects/go/rtmp-service/component/appContext"
	zaplogger "hiholive/shared/go/logger/zap"
	"io"
	"log"
	"net"
)

func main() {
	logger := zaplogger.NewZapLogger(context.Background(), zap.DebugLevel)
	appCtx := appContext.NewAppContextRTPM(logger)
	appCtx.GetLogger()

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1935")
	if err != nil {
		log.Panicf("Failed: %+v", err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Panicf("Failed: %+v", err)
	}

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
	if err := srv.Serve(listener); err != nil {
		log.Panicf("Failed: %+v", err)
	}
}
