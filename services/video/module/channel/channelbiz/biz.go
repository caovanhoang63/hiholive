package channelbiz

import (
	"errors"
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

func (c *channelBiz) Create(ctx context.Context, requester core.Requester, create *channelmodel.ChannelCreate) error {
	channel, err := c.channelRepo.FindUserChannel(ctx, requester.GetUserId())
	if err != nil && !errors.Is(err, core.ErrRecordNotFound) {
		return core.ErrInternalServerError.WithTrace(err)
	}

	if channel != nil {
		return core.ErrBadRequest.WithError("channel already exists")
	}

	create.UserId = requester.GetUserId()

	if err := c.channelRepo.Create(ctx, create); err != nil {
		return core.ErrInternalServerError.WithTrace(err)
	}

	return nil
}
