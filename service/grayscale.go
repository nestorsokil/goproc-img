package service

import (
	"io"

	"github.com/example/goproc-img/util"
)

// RGB2GrayScale converts image file to black & white
func RGB2GrayScale(file io.Reader) (io.Reader, error) {
	grayImage, err := util.Grayscale(file)
	if err != nil {
		return nil, err
	}
	return util.ImageToReader(grayImage)
}
