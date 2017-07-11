package main

import (
	"net/http"
	"log"
	"github.com/julienschmidt/httprouter"
	"github.com/nestorsokil/goproc-img/api-ms/handler"
)

const (
	PING_PATH = "/ping"
	UPLOAD_PATH = "/api/v1/upload"
)

func main() {
	handlers := handler.GetHandlers()
	router := httprouter.New()

	router.GET(PING_PATH, handlers.DoPong)
	router.GET(UPLOAD_PATH, handlers.StoreImageByUrl)
	router.POST(UPLOAD_PATH, handlers.StoreImageByPostData)

	log.Println("Starting server on 8080.")
	log.Fatal(http.ListenAndServe(":8080", router))
}