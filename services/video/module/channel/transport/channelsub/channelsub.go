package channelsub

import (
	"errors"
	"github.com/caovanhoang63/hiholive/services/video/module/channel/channelbiz"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/pubsub"
	"github.com/caovanhoang63/hiholive/shared/golang/subengine"
	"golang.org/x/net/context"
)

type ChannelSub struct {
	biz        channelbiz.ChannelBiz
	serviceCtx srvctx.ServiceContext
}

func NewChannelSub(biz channelbiz.ChannelBiz, serviceCtx srvctx.ServiceContext) *ChannelSub {
	return &ChannelSub{
		biz:        biz,
		serviceCtx: serviceCtx,
	}
}

func (s *ChannelSub) UpdateChannelName() subengine.ConsumerJob {
	return subengine.ConsumerJob{
		Title: "Update channel name",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			data, ok := message.Data.(map[string]interface{})
			if !ok {
				return errors.New("invalid data format")
			}
			id, ok := data["id"].(float64)
			userName := data["userName"].(string)
			displayName := data["displayName"].(string)
			if err := s.biz.UpdateChannelName(ctx, int(id), userName, displayName); err != nil {
				return err
			}
			return nil
		},
	}
}
