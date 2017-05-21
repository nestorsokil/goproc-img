package main

import (
	"fmt"
	"net/http"
	"io"
	"net/url"
)

func GrayScaleHandler(writer http.ResponseWriter, request *http.Request)  {
	file, err := extractImageOrFail(request)
	if err != nil {
		respond400(writer, err.Error())
	} else if res, err := RGB2GrayScale(file); err == nil {
		respond200(writer, res)
	} else {
		respond500(writer, err.Error())
	}
}

func respond400(writer http.ResponseWriter, msg string) {
	writer.WriteHeader(http.StatusBadRequest)
	writer.Header().Set("Content-Type", "text/plain")
	io.WriteString(writer, msg)
}

func respond500(writer http.ResponseWriter, msg string) {
	fmt.Printf("[ERROR] %s\n", msg)
	writer.WriteHeader(http.StatusInternalServerError)
	writer.Header().Set("Content-Type", "text/plain")
	io.WriteString(writer, msg)
}

func respond200(writer http.ResponseWriter, file io.Reader) {
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "image/png")
	io.Copy(writer, file)
}

type RequestError struct {
	desc string
}

func (re RequestError) Error() string {
	return re.desc
}

func extractImageOrFail(request * http.Request) (io.Reader, error) {
	switch request.Method {
	case "GET":
		rq := request.URL.RawQuery
		if values, err := url.ParseQuery(rq); err != nil {
			return nil, RequestError{"Error parsing query string."}
		} else {
			if imageURL, ok := values["image"]; ok {
				imageResponse, err := http.Get(imageURL[0])
				if err != nil {
					return nil, RequestError{"Unable to load image by URL."}
				}
				return imageResponse.Body, nil
			} else {
				return nil, RequestError{"Param 'image' is required."}
			}
		}
	case "POST":
		request.ParseMultipartForm(50 << 20) // 50 mb
		f, _, err := request.FormFile("file")
		return f, err
	default:
		return nil, RequestError{"Bad Request"}
	}
}