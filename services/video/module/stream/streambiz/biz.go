package streambiz

import (
	"fmt"
	"github.com/caovanhoang63/hiholive/services/video/module/channel/channelmodel"
	"github.com/caovanhoang63/hiholive/services/video/module/stream/streammodel"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type StreamRepo interface {
	Create(ctx context.Context, create *streammodel.StreamCreate) error
}

type StreamBiz interface {
	Create(ctx context.Context, requester core.Requester, create *streammodel.StreamCreate) (*streammodel.StreamCreateResponse, error)
}

type ChannelRepo interface {
	FindUserChannel(ctx context.Context, userId int) (*channelmodel.Channel, error)
}

type streamBiz struct {
	streamRepo  StreamRepo
	channelRepo ChannelRepo
}

func NewStreamBiz(streamRepo StreamRepo, channelRepo ChannelRepo) *streamBiz {
	return &streamBiz{
		streamRepo:  streamRepo,
		channelRepo: channelRepo,
	}
}

func (s *streamBiz) Create(ctx context.Context, requester core.Requester, create *streammodel.StreamCreate) (*streammodel.StreamCreateResponse, error) {
	if requester.GetRole() != "streamer" {
		return nil, core.ErrForbidden
	}

	channel, err := s.channelRepo.FindUserChannel(ctx, requester.GetUserId())
	if err != nil {
		return nil, core.ErrInternalServerError.WithTrace(err)
	}

	create.ChannelId = channel.Id
	streamKey, _ := uuid.NewUUID()
	fmt.Println(streamKey)
	create.StreamKey = &streamKey

	if err = s.streamRepo.Create(ctx, create); err != nil {
		fmt.Println(err)
		return nil, core.ErrInternalServerError.WithTrace(err)
	}

	return &streammodel.StreamCreateResponse{
		StreamKey: create.StreamKey,
		RtmpLink:  "http:",
	}, nil

}
