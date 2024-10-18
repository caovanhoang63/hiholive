package main

import (
	"context"
	"github.com/yutopp/go-rtmp"
	"go.uber.org/zap"
	"hiholive/projects/go/rtmp/component/appContext"
	logger "hiholive/shared/go/logger"
	zaplogger "hiholive/shared/go/logger/zap"
	"io"
	"net"
)

func main() {
	// Setup dependencies
	l := zaplogger.NewZapLogger(context.Background(), zap.DebugLevel)
	appCtx := appContext.NewAppContextRTPM(l)
	log := appCtx.GetLogger()
	// Setup dependencies

	// Setup TCP connection
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1935")
	if err != nil {
		log.FatalWithFields("Failed: %+v", logger.Field{"Error": err})
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.FatalWithFields("Failed: %+v", logger.Field{"Error": err})
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
		log.FatalWithFields("Failed: %+v", logger.Field{"Error": err})
	}
}
