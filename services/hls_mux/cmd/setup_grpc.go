package cmd

import (
	"fmt"
	"github.com/caovanhoang63/hiholive/services/hls_mux/component/ffmpegc"
	"github.com/caovanhoang63/hiholive/services/hls_mux/composer"
	"github.com/caovanhoang63/hiholive/shared/golang/core"
	"github.com/caovanhoang63/hiholive/shared/golang/proto/pb"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx"
	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"net"
)

func StartGRPCServices(ffmpeg *ffmpegc.Ffmpeg, serviceCtx srvctx.ServiceContext) {
	configComp := serviceCtx.MustGet(core.KeyCompConf).(core.Config)
	logger := serviceCtx.Logger("grpc")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", configComp.GetGRPCPort()))

	if err != nil {
		log.Fatal(err)
	}

	logger.Infof("GRPC Server is listening on %d ...\n", configComp.GetGRPCPort())

	s := grpc.NewServer()

	pb.RegisterHlsServiceServer(s, composer.ComposeHlsGRPCService(ffmpeg, serviceCtx))

	if err = s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
