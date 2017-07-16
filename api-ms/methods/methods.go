package methods

import (
	"github.com/nestorsokil/goproc-img/api-ms/config"
	"github.com/nestorsokil/goproc-img/api-ms/service"
	"github.com/nestorsokil/goproc-img/api-ms/util"
	"io"
	"net/http"
)

const (
	GET_FILE_URL  = "fileUrl"
	GET_FILE_NAME = "fileName"

	POST_FILE_DATA = "fileData"
	POST_FILE_NAME = "fileName"
)

type Handlers struct {
	config config.ApiConfig
}

func (h Handlers) StoreFileByPostData(
	response http.ResponseWriter, request *http.Request) {

	request.ParseMultipartForm(50 << 20) // 50 mb
	data, _, err := request.FormFile(POST_FILE_DATA)
	if err != nil {
		util.Respond400(response, "Could not get file from POST.")
		return
	}
	name := request.FormValue(POST_FILE_NAME)
	client, err := util.GetClient(h.config.StoreMsUrl)
	if err != nil {
		util.Respond400(response, err.Error())
		return
	}
	url, err := service.SaveImage(name, data, client)
	if err != nil {
		util.Respond400(response, err.Error())
		return
	}
	io.WriteString(response, url)
}

func (h Handlers) StoreFileByUrl(
	response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	imageUrl := query.Get(GET_FILE_URL)
	imageName := query.Get(GET_FILE_NAME)
	imageResponse, err := http.Get(imageUrl)
	client, err := util.GetClient(h.config.StoreMsUrl)
	if err != nil {
		util.Respond400(response, err.Error())
		return
	}
	if err != nil {
		util.Respond400(response, "Unable to load image by URL.")
		return
	}
	url, err := service.SaveImage(imageName, imageResponse.Body, client)
	if err != nil {
		util.Respond400(response, err.Error())
		return
	}
	io.WriteString(response, url)
}

func (h Handlers) DoPong(response http.ResponseWriter, _ *http.Request) {
	io.WriteString(response, "Pong!")
}

func GetHandlers() Handlers {
	return Handlers{config: config.LoadConfig()}
}
