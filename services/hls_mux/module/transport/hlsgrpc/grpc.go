package hlsgrpc

import (
	"fmt"
	"github.com/caovanhoang63/hiholive/services/hls_mux/component/ffmpegc"
	"github.com/caovanhoang63/hiholive/shared/go/proto/pb"
	"golang.org/x/net/context"
	"time"
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
	fps := req.Fps
	resolution := req.Resolution
	streamId := req.StreamId

	go func() {
		fn := h.Ffmpeg.NewStream(streamId, link, key, int(fps), int(resolution))

		go func() {
			time.Sleep(20 * time.Second)

			for i := 0; i < 10; i++ {
				fmt.Println(i)
			}
			fn()
		}()

	}()

	return &pb.NewHlsStreamResp{}, nil
}
