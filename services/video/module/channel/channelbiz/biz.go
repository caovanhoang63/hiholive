package channelbiz

import (
	"github.com/caovanhoang63/hiholive/services/video/module/channel/channelmodel"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"golang.org/x/net/context"
)

type ChannelRepo interface {
	Create(ctx context.Context, create *channelmodel.ChannelCreate) error
	FindUserChannel(ctx context.Context, userId int) (*channelmodel.Channel, error)
}

type ChannelBiz interface {
	Create(ctx context.Context, requester core.Requester, create *channelmodel.ChannelCreate) error
}

type channelBiz struct {
	channelRepo ChannelRepo
}

func NewChannelBiz(channelRepo ChannelRepo) *channelBiz {
	return &channelBiz{channelRepo: channelRepo}
}

func (c channelBiz) Create(ctx context.Context, requester core.Requester, create *channelmodel.ChannelCreate) error {
}
