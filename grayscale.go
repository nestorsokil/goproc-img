package main

import (
	"io"
)

// RGB2GrayScale converts image file to black & white
func RGB2GrayScale(file io.Reader) (io.Reader, error) {
	grayImage, err := Grayscale(file)
	if err != nil {
		return nil, err
	}
	return ImageToReader(grayImage)
}
