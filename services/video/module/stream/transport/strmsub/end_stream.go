package strmsub

import (
	"github.com/caovanhoang63/hiholive/services/video/module/stream/streammodel"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/caovanhoang63/hiholive/shared/go/srvctx/components/pubsub"
	"github.com/caovanhoang63/hiholive/shared/go/subengine"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

func (s *StreamSub) EndStream() subengine.ConsumerJob {
	return subengine.ConsumerJob{
		Title: "Update stream's state to ended",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			data, ok := message.Data.(map[string]any)
			if !ok {
				return errors.New("Invalid data format")
			}

			uid, ok := data["stream_id"].(string)

			if !ok {
				return errors.New("Invalid data format")
			}

			id, err := core.FromBase58(uid)
			if err != nil {
				return err
			}
			if err = s.biz.UpdateStreamState(ctx, nil, int(id.GetLocalID()), streammodel.StreamStateEnded); err != nil {
				return err
			}
			return nil
		},
	}
}
