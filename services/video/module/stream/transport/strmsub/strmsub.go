package strmsub

import (
	"github.com/caovanhoang63/hiholive/services/video/module/stream/streambiz"
	"github.com/caovanhoang63/hiholive/shared/golang/core"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/pubsub"
	"github.com/caovanhoang63/hiholive/shared/golang/subengine"
	"golang.org/x/net/context"
)

type StreamSub struct {
	biz        streambiz.StreamBiz
	serviceCtx srvctx.ServiceContext
}

func NewStreamSub(biz streambiz.StreamBiz, serviceCtx srvctx.ServiceContext) *StreamSub {
	return &StreamSub{
		biz:        biz,
		serviceCtx: serviceCtx,
	}
}

func (s *StreamSub) StartStream() subengine.ConsumerJob {
	return subengine.ConsumerJob{
		Title: "Update stream's state to running",
		Handler: func(ctx context.Context, message *pubsub.Message) error {

			if uid, ok := message.Data.(string); ok {
				id, err := core.FromBase58(uid)
				if err != nil {
					return err
				}

				if err = s.biz.UpdateStreamState(ctx, nil, int(id.GetLocalID()), "running"); err != nil {
					return err
				}
			}
			return nil
		},
	}
}
