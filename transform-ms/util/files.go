package util

import (
	"bytes"
	"image"
	"image/png"
	"io"
)

// ImageToReader converts in-memory image to an io.Reader
func ImageToReader(img image.Image) (io.Reader, error) {
	buf := new(bytes.Buffer)
	err := png.Encode(buf, img)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(buf.Bytes()), nil
}
