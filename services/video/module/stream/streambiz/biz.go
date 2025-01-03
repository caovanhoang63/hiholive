package streambiz

import (
	"fmt"
	"github.com/caovanhoang63/hiholive/services/video/module/channel/channelmodel"
	"github.com/caovanhoang63/hiholive/services/video/module/stream/streammodel"
	"github.com/caovanhoang63/hiholive/shared/golang/core"
	"github.com/caovanhoang63/hiholive/shared/golang/srvctx/components/pubsub"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type StreamRepo interface {
	Create(ctx context.Context, create *streammodel.StreamCreate) error
	FindStreamByID(ctx context.Context, id int) (*streammodel.Stream, error)
	UpdateStream(ctx context.Context, id int, update *streammodel.StreamUpdate) error
	FindStreams(ctx context.Context, filter *streammodel.StreamFilter, paging *core.Paging) ([]streammodel.StreamList, error)
	UpdateStreamView(ctx context.Context, id, view int) error
}

type StreamBiz interface {
	Create(ctx context.Context, requester core.Requester, create *streammodel.StreamCreate) (*streammodel.StreamCreateResponse, error)
	FindStreamById(ctx context.Context, id int) (*streammodel.Stream, error)
	UpdateStreamState(ctx context.Context, requester core.Requester, id int, state string) error
	FindStreams(ctx context.Context, filter *streammodel.StreamFilter, paging *core.Paging) ([]streammodel.StreamList, error)
	UpdateStreamView(ctx context.Context, id, view int) error
}

type ChannelRepo interface {
	FindUserChannel(ctx context.Context, userId int) (*channelmodel.Channel, error)
}

type streamBiz struct {
	streamRepo  StreamRepo
	channelRepo ChannelRepo
	ps          pubsub.Pubsub
}

func NewStreamBiz(streamRepo StreamRepo, channelRepo ChannelRepo, ps pubsub.Pubsub) *streamBiz {
	return &streamBiz{
		streamRepo:  streamRepo,
		channelRepo: channelRepo,
		ps:          ps,
	}
}

func (s *streamBiz) UpdateStreamView(ctx context.Context, id, view int) error {
	if err := s.streamRepo.UpdateStreamView(ctx, id, view); err != nil {
		return core.ErrInternalServerError.WithWrap(err)
	}
	return nil
}
func (s *streamBiz) FindStreams(ctx context.Context, filter *streammodel.StreamFilter, paging *core.Paging) ([]streammodel.StreamList, error) {
	if err := filter.Process(); err != nil {
		return nil, core.ErrBadRequest.WithError(err.Error())
	}

	data, err := s.streamRepo.FindStreams(ctx, filter, paging)
	if err != nil {
		return nil, core.ErrInternalServerError.WithWrap(err)
	}
	return data, nil
}

func (s *streamBiz) UpdateStreamState(ctx context.Context, requester core.Requester, id int, state string) error {
	stream, err := s.FindStreamById(ctx, id)
	if err != nil {
		return core.ErrInternalServerError.WithWrap(err)
	}

	if stream == nil {
		return core.ErrNotFound
	}

	if stream.State == state {
		return nil
	}

	if err = s.streamRepo.UpdateStream(ctx, id, &streammodel.StreamUpdate{
		State: &state,
	}); err != nil {
		return core.ErrInternalServerError.WithWrap(err)
	}
	return nil
}

func (s *streamBiz) FindStreamById(ctx context.Context, id int) (*streammodel.Stream, error) {

	r, err := s.streamRepo.FindStreamByID(ctx, id)
	if err != nil {
		return nil, core.ErrInternalServerError.WithWrap(err)
	}

	if r == nil {
		return nil, core.ErrNotFound
	}

	return r, nil
}

func (s *streamBiz) Create(ctx context.Context, requester core.Requester, create *streammodel.StreamCreate) (*streammodel.StreamCreateResponse, error) {
	if requester.GetRole() != "streamer" {
		return nil, core.ErrForbidden
	}

	channel, err := s.channelRepo.FindUserChannel(ctx, requester.GetUserId())
	if err != nil {
		return nil, core.ErrInternalServerError.WithWrap(err)
	}

	create.ChannelId = channel.Id
	streamKey, _ := uuid.NewUUID()
	create.StreamKey = &streamKey
	if err = s.streamRepo.Create(ctx, create); err != nil {
		fmt.Println(err)
		return nil, core.ErrInternalServerError.WithTrace(err)
	}

	_ = s.ps.Publish(ctx, core.TopicStreamCreate, pubsub.NewMessage(create))

	return &streammodel.StreamCreateResponse{
		StreamId:  create.Uid,
		StreamKey: create.StreamKey,
		RtmpLink:  "rtmp://rtmp.hiholive.fun/stream/",
	}, nil

}
