package service

import (
	"io"
	"bytes"
	"github.com/nestorsokil/goproc-img/store-ms/api"
	"golang.org/x/net/context"
)

func SaveImage(filename string, content io.Reader, client api.StoreServiceClient)(url string, err error){

	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(content)
	data := buf.Bytes()
	req := &api.SaveRequest{Filename: filename, Data: data}
	result, err := client.SaveFile(context.Background(), req)
	if err != nil  {
		return "", err
	}
	return result.URL, nil

}