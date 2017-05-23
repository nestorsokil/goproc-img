package main

import (
	"bytes"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"io"
)

func RGB2GrayScale(file io.Reader) (io.Reader, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	gray := image.NewGray(bounds)
	xMax, yMax := bounds.Max.X, bounds.Max.Y

	for x := 0; x < xMax; x++ {
		for y := 0; y < yMax; y++ {
			rgba := img.At(x, y)
			gray.Set(x, y, rgba)
		}
	}

	buf := new(bytes.Buffer)
	if err := png.Encode(buf, gray); err == nil {
		return bytes.NewReader(buf.Bytes()), nil
	} else {
		return nil, err
	}
}
