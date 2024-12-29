package composer

import (
	"github.com/caovanhoang63/hiholive/services/hls_mux/component/ffmpegc"
	"github.com/caovanhoang63/hiholive/services/hls_mux/module/transport/hlsgrpc"
	"github.com/caovanhoang63/hiholive/shared/golang/proto/pb"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx"
)

func ComposeHlsGRPCService(ffmpeg *ffmpegc.Ffmpeg, serviceCtx srvctx.ServiceContext) pb.HlsServiceServer {
	return hlsgrpc.NewHlsGRPC(ffmpeg)
}
