package service

import (
	"image"
	"io"

	"github.com/nestorsokil/goproc-img/transform-ms/util"
	"github.com/nfnt/resize"
)

func Resize(imageFile io.Reader, sizeX, sizeY uint) (io.Reader, error) {
	im, _, err := image.Decode(imageFile)
	if err != nil {
		return nil, err
	}

	resized := resize.Resize(sizeX, sizeY, im, resize.Bicubic)
	return util.ImageToReader(resized)
}
