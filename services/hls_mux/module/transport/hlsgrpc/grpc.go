package hlsgrpc

import (
	"fmt"
	"github.com/caovanhoang63/hiholive/services/hls_mux/component/ffmpegc"
	"github.com/caovanhoang63/hiholive/shared/go/proto/pb"
	"golang.org/x/net/context"
)

type hlsGRPC struct {
	Ffmpeg *ffmpegc.Ffmpeg
}

func NewHlsGRPC(Ffmpeg *ffmpegc.Ffmpeg) *hlsGRPC {
	return &hlsGRPC{Ffmpeg: Ffmpeg}
}

func (h *hlsGRPC) NewHlsStream(ctx context.Context, req *pb.NewHlsStreamReq) (*pb.NewHlsStreamResp, error) {
	key := req.StreamKey
	link := req.ServerUrl
	fmt.Println("Calllll")

	go h.Ffmpeg.NewStream(link, key)

	return &pb.NewHlsStreamResp{}, nil
}
