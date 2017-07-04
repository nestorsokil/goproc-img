package main

import (
	"net/http"
	"io"
	"log"

	"github.com/julienschmidt/httprouter"
)

const (
	PING_PATH = "/ping"
	UPLOAD_PATH = "/api/v1/upload"
	IMAGE_URL = "imageUrl"
)

func main() {
	router := httprouter.New()

	router.GET(PING_PATH, doPong)
	router.GET(UPLOAD_PATH, doGet)
	router.POST(UPLOAD_PATH, doPost)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func doPost(response http.ResponseWriter, request *http.Request,  _ httprouter.Params) {
	request.ParseMultipartForm(50 << 20) // 50 mb
	f, _, err := request.FormFile("file")
	if err != nil {
		respond400(response, "Could not get file from POST.")
		return
	}

	saveImage(f)
}

func doGet(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	query := request.URL.Query()
	imageUrl := query.Get(IMAGE_URL)
	imageResponse, err := http.Get(imageUrl)
	if err != nil {
		respond400(response, "Unable to load image by URL.")
	}
	saveImage(imageResponse.Body)
}

func doPong(response http.ResponseWriter, _ *http.Request,  _ httprouter.Params) {
	io.WriteString(response, "Pong!")
}

func saveImage(content io.Reader) {

}


func respond400(writer http.ResponseWriter, msg string) {
	writer.WriteHeader(http.StatusBadRequest)
	writer.Header().Set("Content-Type", "text/plain")
	io.WriteString(writer, msg)
}