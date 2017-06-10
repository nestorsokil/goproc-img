package service

import (
	"image"
	"io"

	"github.com/nestorsokil/goproc-img/util"
	"github.com/nfnt/resize"
)

func Resize(imagefile io.Reader, sizex, sizey uint) (io.Reader, error) {
	im, _, err := image.Decode(imagefile)
	if err != nil {
		return nil, err
	}

	resized := resize.Resize(sizex, sizey, im, resize.Bicubic)
	return util.ImageToReader(resized)
}
