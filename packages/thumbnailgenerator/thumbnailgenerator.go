package thumbnailgenerator

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"shopping-lists-and-recipes/packages/shared"

	"github.com/disintegration/imaging"
)

var (
	ErrDstNil = errors.New("graphics: dst is nil")
	ErrSrcNil = errors.New("graphics: src is nil")
)

type Dimensions struct {
	Width  int
	Height int
}

func Process(f io.ReadSeeker, filetype string, dims Dimensions) (image.Image, error) {

	f.Seek(0, io.SeekStart)

	var result image.Image
	var base image.Image
	var err error

	if filetype == "image/jpeg" || filetype == "image/jpg" {
		base, err = jpeg.Decode(f)
	} else if filetype == "image/png" {
		base, err = png.Decode(f)
	} else if filetype == "image/gif" {
		base, err = gif.Decode(f)
	} else {
		return result, shared.ErrUnsupportedFileType
	}

	if err != nil {
		return result, err
	}

	result = imaging.Thumbnail(base, dims.Width, dims.Height, imaging.Lanczos)

	return result, nil
}
