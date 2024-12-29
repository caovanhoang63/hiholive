package uploadprovider

import (
	"context"
	"github.com/caovanhoang63/hiholive/shared/golang/core"
)

type UploadProvider interface {
	SaveImageUploaded(ctx context.Context, data []byte, dst string) (*core.Image, error)
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*core.File, error)
}
