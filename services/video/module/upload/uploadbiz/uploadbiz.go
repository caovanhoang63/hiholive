package uploadbiz

import (
	"bytes"
	"fmt"
	"github.com/caovanhoang63/hiholive/shared/go/core"
	"github.com/caovanhoang63/hiholive/shared/go/uploadprovider"
	"golang.org/x/net/context"
	"image"
	"io"
	"path/filepath"
	"strings"
	"time"
)

type UploadBiz interface {
	UploadImage(ctx context.Context, data []byte, folder, fileName string) (*core.Image, error)
}

type uploadBiz struct {
	provider uploadprovider.UploadProvider
}

func NewUploadBiz(provider uploadprovider.UploadProvider) *uploadBiz {
	return &uploadBiz{provider: provider}
}

func (biz *uploadBiz) UploadImage(ctx context.Context, data []byte, folder, fileName string) (*core.Image, error) {
	fileBytes := bytes.NewBuffer(data)
	w, h, err := getImageDimension(fileBytes)
	if err != nil {
		return nil, core.ErrInvalidInput("file")
	}
	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}
	fileExt := filepath.Ext(fileName)
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt)
	img, err := biz.provider.SaveImageUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))
	if err != nil {
		return nil, core.ErrInternalServerError.WithWrap(err)
	}

	img.Width = w
	img.Height = h
	img.Extension = fileExt

	return img, nil
}

// getImageDimension returns the width and height of an image file
func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)

	if err != nil {
		fmt.Println("err", err)
		return 0, 0, err
	}

	return img.Width, img.Height, err
}
