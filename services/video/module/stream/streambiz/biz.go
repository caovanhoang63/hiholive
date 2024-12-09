package streambiz

import (
	"github.com/caovanhoang63/hiholive/services/video/module/stream/streammodel"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"golang.org/x/net/context"
)

type StreamRepo interface {
	Create(ctx context.Context, create streammodel.StreamCreate) error
}

type StreamBiz interface {
	Create(ctx context.Context, requester core.Requester, create streammodel.StreamCreate) error
}

type streamRepo struct {
	repo StreamRepo
}

func (s *streamRepo) Create(ctx context.Context, requester core.Requester, create streammodel.StreamCreate) error {
	return s.repo.Create(ctx, create)
}
