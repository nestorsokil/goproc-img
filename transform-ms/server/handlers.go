package server

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/nestorsokil/goproc-img/transform-ms/service"
)

type handler func(file io.Reader) (io.Reader, error)

// GrayScaleHandler is a REST handle func for grayscale convertion
func GrayScaleHandler(writer http.ResponseWriter, request *http.Request) {
	handleGeneric(writer, request, service.RGB2GrayScale)
}

// BinaryHandler is a REST handle func for binary convertion
func BinaryHandler(writer http.ResponseWriter, request *http.Request) {
	handleGeneric(writer, request, service.RGB2Binary)
}

// NegativeHandler is a REST handle func for negative convertion
func NegativeHandler(writer http.ResponseWriter, request *http.Request) {
	handleGeneric(writer, request, service.RGB2Negative)
}

func ResizeHandler(writer http.ResponseWriter, request *http.Request) {
	rq := request.URL.RawQuery
	values, err := url.ParseQuery(rq)
	widthStr := values["width"][0]
	heightStr := values["height"][0]
	width, _ := strconv.ParseUint(widthStr, 10, 32)
	height, _ := strconv.ParseUint(heightStr, 10, 32)

	file, err := extractImageOrFail(request)
	if err != nil {
		respond400(writer, err.Error())
	} else if res, err := service.Resize(file, uint(width), uint(height)); err == nil {
		respond200(writer, res)
	} else {
		respond500(writer, err.Error())
	}
}

func handleGeneric(writer http.ResponseWriter, request *http.Request,
	handleFunc handler) {
	file, err := extractImageOrFail(request)
	if err != nil {
		respond400(writer, err.Error())
	} else if res, err := handleFunc(file); err == nil {
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

type requestError struct {
	desc string
}

func (re requestError) Error() string {
	return re.desc
}

func extractImageOrFail(request *http.Request) (io.Reader, error) {
	switch request.Method {
	case "GET":
		rq := request.URL.RawQuery
		values, err := url.ParseQuery(rq)
		if err != nil {
			return nil, requestError{"Error parsing query string."}
		}
		if imageURL, ok := values["image"]; ok {
			imageResponse, err := http.Get(imageURL[0])
			if err != nil {
				return nil, requestError{"Unable to load image by URL."}
			}
			return imageResponse.Body, nil
		}
		return nil, requestError{"Param 'image' is required."}

	case "POST":
		request.ParseMultipartForm(50 << 20) // 50 mb
		f, _, err := request.FormFile("file")
		return f, err

	default:
		return nil, requestError{"Bad Request"}
	}
}
