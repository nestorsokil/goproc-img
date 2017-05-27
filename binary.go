package main

import (
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	"io"
)

func RGB2Binary(file io.Reader) (io.Reader, error) {
	grayImage, err := Grayscale(file)
	if err != nil {
		return nil, err
	}
	threshold := OtsuThreshold(grayImage)

	bounds := grayImage.Bounds()
	binImage := image.NewRGBA(bounds)
	xMax, yMax := bounds.Max.X, bounds.Max.Y

	for x := 0; x < xMax; x++ {
		for y := 0; y < yMax; y++ {
			r, _, _, _ := grayImage.At(x, y).RGBA()
			var col color.Color
			if uint8(r) > threshold {
				col = color.White
			} else {
				col = color.Black
			}
			binImage.Set(x, y, col)
		}
	}

	return ImageToReader(binImage)
}
