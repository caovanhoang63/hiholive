package channelbiz

import (
	"errors"
	"github.com/caovanhoang63/hiholive/services/video/common"
	"github.com/caovanhoang63/hiholive/services/video/module/channel/channelmodel"
	"github.com/caovanhoang63/hiholive/shared/golang/core"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/pubsub"
	"golang.org/x/net/context"
	"strings"
)

type ChannelRepo interface {
	Create(ctx context.Context, create *channelmodel.ChannelCreate) error
	FindUserChannel(ctx context.Context, userId int) (*channelmodel.Channel, error)
	FindChannelById(ctx context.Context, channelId int) (*channelmodel.Channel, error)
	FindChannelByUserName(ctx context.Context, userName string) (*channelmodel.Channel, error)
	FindChannels(ctx context.Context, filter *channelmodel.ChannelFilter, paging *core.Paging) ([]channelmodel.Channel, error)
}

type ChannelBiz interface {
	Create(ctx context.Context, requester core.Requester, create *channelmodel.ChannelCreate) error
	FindUserChannel(ctx context.Context, userId int) (*channelmodel.Channel, error)
	FindChannelById(ctx context.Context, channelId int) (*channelmodel.Channel, error)
	FindChannelByUserName(ctx context.Context, userName string) (*channelmodel.Channel, error)
	FindChannels(ctx context.Context, filter *channelmodel.ChannelFilter, paging *core.Paging) ([]channelmodel.Channel, error)
}

type UserRepo interface {
	GetUserById(ctx context.Context, id int) (*common.User, error)
}

type channelBiz struct {
	channelRepo ChannelRepo
	userRepo    UserRepo
	ps          pubsub.Pubsub
}

func NewChannelBiz(channelRepo ChannelRepo, userRepo UserRepo, ps pubsub.Pubsub) *channelBiz {
	return &channelBiz{channelRepo: channelRepo, ps: ps,
		userRepo: userRepo}
}

func (c *channelBiz) FindChannelByUserName(ctx context.Context, userName string) (*channelmodel.Channel, error) {
	channel, err := c.channelRepo.FindChannelByUserName(ctx, userName)
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			return nil, core.ErrNotFound
		}
		return nil, core.ErrInternalServerError.WithWrap(err)
	}
	return channel, nil
}

func (c *channelBiz) Create(ctx context.Context, requester core.Requester, create *channelmodel.ChannelCreate) error {
	if field, err := core.Validator.ValidateField(create); err != nil {
		return core.ErrInvalidInput(field)
	}

	channel, err := c.channelRepo.FindUserChannel(ctx, requester.GetUserId())

	if err != nil && !errors.Is(err, core.ErrRecordNotFound) {
		return core.ErrInternalServerError.WithWrap(err)
	}

	if channel != nil {
		return core.ErrBadRequest.WithError("channel already exists")
	}
	user, err := c.userRepo.GetUserById(ctx, requester.GetUserId())

	if err != nil || user == nil {
		return core.ErrInternalServerError.WithWrap(err)
	}

	if strings.ToLower(create.UserName) != create.UserName {
		return core.ErrInvalidInput("userName")
	}

	if strings.ToLower(create.UserName) != strings.ToLower(create.DisplayName) {
		return core.ErrInvalidInput("displayName")
	}

	if user.DisplayName != create.DisplayName || user.UserName != create.UserName {
		// TODO : UPDATE userName and DisplayName
	}

	create.UserId = requester.GetUserId()

	if err = c.channelRepo.Create(ctx, create); err != nil {
		return core.ErrInternalServerError.WithWrap(err)
	}

	_ = c.ps.Publish(ctx, core.TopicCreateChannel, pubsub.NewMessage(create.UserId))
	return nil
}

func (c *channelBiz) FindUserChannel(ctx context.Context, userId int) (*channelmodel.Channel, error) {
	channel, err := c.channelRepo.FindUserChannel(ctx, userId)
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			return nil, core.ErrNotFound
		}
		return nil, core.ErrInternalServerError.WithWrap(err)
	}
	return channel, nil
}

func (c *channelBiz) FindChannelById(ctx context.Context, channelId int) (*channelmodel.Channel, error) {
	channel, err := c.channelRepo.FindChannelById(ctx, channelId)
	if err != nil {
		if errors.Is(err, core.ErrRecordNotFound) {
			return nil, core.ErrNotFound
		}
		return nil, core.ErrInternalServerError.WithWrap(err)
	}
	return channel, nil
}

func (c *channelBiz) FindChannels(ctx context.Context, filter *channelmodel.ChannelFilter, paging *core.Paging) ([]channelmodel.Channel, error) {
	channels, err := c.channelRepo.FindChannels(ctx, filter, paging)
	if err != nil {
		return nil, core.ErrInternalServerError.WithWrap(err)
	}
	return channels, nil
}
