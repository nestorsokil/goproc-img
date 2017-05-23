package main

import (
	"bytes"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
)

func RGB2Binary(file io.Reader) (io.Reader, error) {
	grayFile, _ := RGB2GrayScale(file)
	asBytes, err := ioutil.ReadAll(grayFile)
	gray, _, err := image.Decode(bytes.NewReader(asBytes))

	if err != nil {
		return nil, err
	}
	threshold := uint32(otsuThreshold(asBytes))

	bounds := gray.Bounds()
	binImage := image.NewRGBA(bounds)
	xMax, yMax := bounds.Max.X, bounds.Max.Y

	for x := 0; x < xMax; x++ {
		for y := 0; y < yMax; y++ {
			r, g, b, _ := gray.At(x, y).RGBA()
			var col color.Color
			if r + g + b > threshold {
					col = color.Black
			} else {
				col = color.White
			}
			binImage.Set(x, y, col)
		}
	}

	buf := new(bytes.Buffer)
	if err := png.Encode(buf, binImage); err == nil {
		return bytes.NewReader(buf.Bytes()), nil
	} else {
		return nil, err
	}
}

func otsuThreshold(img []byte) int {
	total := len(img)
	histData := make(map[int]int)
	ptr := 0
	for ptr < len(img) {
		h := 0xFF & img[ptr]
		histData[int(h)]++
		ptr++
	}

	sum := 0
	for i := 0; i < 256; i++ {
		sum += i * histData[i]
	}

	var sumB, wB, wF, max, threshold int
	for i := 0; i < 256; i++ {
		wB += histData[i]
		if wB != 0 {
			wF = total - wB
			if wF == 0 {
				break
			}
			sumB += i * histData[i]
			mB := sumB / wB
			mF := (sum - sumB) / wF
			between := wB * wF * (mB - mF) * (mB - mF)
			if between > max {
				max = between
				threshold = i
			}
		}
	}
	return threshold
}
