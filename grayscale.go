package main

import (
	"io"
)

func RGB2GrayScale(file io.Reader) (io.Reader, error) {
	grayImage, err := Grayscale(file)
	if err != nil {
		return nil, err
	}
	return ImageToReader(grayImage)
}
