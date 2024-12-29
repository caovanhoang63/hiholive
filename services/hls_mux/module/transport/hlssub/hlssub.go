package hlssub

import (
	"github.com/caovanhoang63/hiholive/services/hls_mux/component/ffmpegc"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/pubsub"
	"github.com/caovanhoang63/hiholive/shared/golang/subengine"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type HLSSub struct {
	f *ffmpegc.Ffmpeg
}

func NewHLSSub(f *ffmpegc.Ffmpeg) *HLSSub {
	return &HLSSub{f: f}
}

func (h *HLSSub) OnStopStream() subengine.ConsumerJob {
	return subengine.ConsumerJob{
		Title: "Stop FFMPEG Stream",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			data, ok := message.Data.(map[string]interface{})
			if !ok {
				return errors.New("invalid data")
			}

			id, ok := data["stream_id"].(string)

			if !ok {
				return errors.New("invalid id")
			}
			return h.f.StopStream(id)
		},
	}
}
