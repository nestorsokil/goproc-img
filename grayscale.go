package main

import (
	"io"
	"image"
	"image/png"
	"bytes"
	_ "image/jpeg"
	_ "image/gif"
)

func RGB2GrayScale(file io.Reader) (io.Reader, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	gray := image.NewGray(bounds)
	xmax, ymax := bounds.Max.X, bounds.Max.Y

	/*factor := 2

	dx, dy := xmax/factor, ymax/factor

	xprev, yprev := 0, 0
	for i := 0; i < factor; i++ {
		x := xprev + dx
		if x > xmax {
			x = xmax
		}
		for j := 0; j < factor; j++ {
			y := yprev + dy
			if y > ymax {
				y = ymax
			}
			rect := image.Rectangle{
				Min: image.Point{X: xprev, Y: yprev},
				Max: image.Point{X: x, Y: y},
			}
			sub := image.NewGray(rect)
			go proc(img, sub, rect, *gray)
			yprev = y
		}
		xprev = x
	}*/

	for x := 0; x < xmax; x++ {
		for y := 0; y < ymax; y++ {
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

/*
func proc(img image.Image, gray *image.Gray, rect image.Rectangle, full image.Gray) {
	bounds := gray.Bounds()
	sizeX, sizeY := bounds.Max.X, bounds.Max.Y
	for x := 0; x < sizeX; x++ {
		for y := 0; y < sizeY; y++ {
			rgba := img.At(x, y)
			gray.Set(x, y, rgba)
		}
	}
	draw.Draw(full, rect, gray, image.Point{0,0}, draw.Src)
}
*/
