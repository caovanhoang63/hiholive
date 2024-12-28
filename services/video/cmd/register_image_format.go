package cmd

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
)

// RegisterImageFormat registers the standard library's image formats.
func RegisterImageFormat() {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("gif", "gif", gif.Decode, gif.DecodeConfig)
}
