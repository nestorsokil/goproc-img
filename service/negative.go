package service

import (
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"

	"github.com/nestorsokil/goproc-img/util"
)

// RGB2Negative converts rgb image to negative colors
func RGB2Negative(file io.Reader) (io.Reader, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	imgNeg := image.NewRGBA(bounds)
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			src := img.At(x, y)
			neg := negative(src)
			imgNeg.Set(x, y, neg)
		}
	}
	return util.ImageToReader(imgNeg)
}

func negative(col color.Color) color.Color {
	r, g, b, a := col.RGBA()
	rn, gn, bn, an := uint16(255-r), uint16(255-g), uint16(255-b), uint16(a)
	newCol := color.RGBA64{rn, gn, bn, an}
	return newCol
}
