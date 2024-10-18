package uploadprovider

import (
	"context"
	"hiholive/shared/go/utils"
)

type UploadProvider interface {
	SaveImageUploaded(ctx context.Context, data []byte, dst string) (*utils.Image, error)
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*utils.File, error)
}
