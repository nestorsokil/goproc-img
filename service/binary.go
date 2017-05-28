package service

import (
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	"io"

	"github.com/example/goproc-img/util"
)

// RGB2Binary converts the image file to a binary image
func RGB2Binary(file io.Reader) (io.Reader, error) {
	grayImage, err := util.Grayscale(file)
	if err != nil {
		return nil, err
	}
	threshold := util.OtsuThreshold(grayImage)

	bounds := grayImage.Bounds()
	binImage := image.NewRGBA(bounds)
	xMax, yMax := bounds.Max.X, bounds.Max.Y

	for x := 0; x < xMax; x++ {
		for y := 0; y < yMax; y++ {
			// RGBA() for gray returns pixel intensity
			intensity, _, _, _ := grayImage.At(x, y).RGBA()
			var col color.Color
			if uint8(intensity) > threshold {
				col = color.White
			} else {
				col = color.Black
			}
			binImage.Set(x, y, col)
		}
	}

	return util.ImageToReader(binImage)
}
