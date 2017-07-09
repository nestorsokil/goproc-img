package main

import (
	"net/http"
	"io"
	"log"

	"github.com/julienschmidt/httprouter"
	"github.com/nestorsokil/goproc-img/store-ms/api"
	"bytes"
	"os"
	"encoding/json"
)

const (
	PING_PATH = "/ping"
	UPLOAD_PATH = "/api/v1/upload"
	CONFIG_FILE = "conf.json"

	GET_FILE_URL  = "fileUrl"
	GET_FILE_NAME = "fileName"

	POST_FILE_DATA = "fileData"
	POST_FILE_NAME = "fileName"
)

var config = struct {
	StoreMsUrl string `json:"store_ms_url"`
}{}

func initConfig() {
	configFile, err := os.Open(CONFIG_FILE)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	err = json.NewDecoder(configFile).Decode(&config)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}
}

func doPost(response http.ResponseWriter, request *http.Request,  _ httprouter.Params) {
	request.ParseMultipartForm(50 << 20) // 50 mb
	f, _, err := request.FormFile(POST_FILE_DATA)
	if err != nil {
		respond400(response, "Could not get file from POST.")
		return
	}
	name := request.FormValue(POST_FILE_NAME)

	url, err := saveImage(name, f)
	if err != nil {
		respond400(response, err.Error())
	}
	io.WriteString(response, url)
}

func doGet(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	query := request.URL.Query()
	imageUrl := query.Get(GET_FILE_URL)
	imageName := query.Get(GET_FILE_NAME)
	imageResponse, err := http.Get(imageUrl)
	if err != nil {
		respond400(response, "Unable to load image by URL.")
	}
	url, err := saveImage(imageName, imageResponse.Body)
	if err != nil {
		respond400(response, err.Error())
	}
	io.WriteString(response, url)
}

func doPong(response http.ResponseWriter, _ *http.Request,  _ httprouter.Params) {
	io.WriteString(response, "Pong!")
}

func saveImage(filename string, content io.Reader) (url string, err error){
	client, err := api.NewClient("localhost:8081")
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(content)
	data := buf.Bytes()
	req := &api.SaveRequest{Filename: filename, Data: data}
	result, err := client.SaveFile(req)
	if err != nil  {
		return "", err
	}
	return result.URL, nil
}

func respond400(writer http.ResponseWriter, msg string) {
	writer.WriteHeader(http.StatusBadRequest)
	writer.Header().Set("Content-Type", "text/plain")
	io.WriteString(writer, msg)
}

func main() {
	initConfig()
	router := httprouter.New()

	router.GET(PING_PATH, doPong)
	router.GET(UPLOAD_PATH, doGet)
	router.POST(UPLOAD_PATH, doPost)

	log.Fatal(http.ListenAndServe(":8080", router))
}