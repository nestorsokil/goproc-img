package main

import (

	"github.com/nestorsokil/goproc-img/api-ms/methods"
	"log"
	"net/http"
	"os"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

const (
	PING_PATH   = "/ping"
	UPLOAD_PATH = "/api/v1/upload"
)

func main() {
	file, err := os.OpenFile("api.log",
		os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)

	h := methods.GetHandlers()
	router := mux.NewRouter()

	router.Handle(PING_PATH, http.HandlerFunc(h.DoPong)).Methods("GET")
	router.Handle(UPLOAD_PATH, http.HandlerFunc(h.StoreFileByUrl)).Methods("GET")
	router.Handle(UPLOAD_PATH, http.HandlerFunc(h.StoreFileByPostData)).Methods("POST")

	withLogging := handlers.LoggingHandler(file, router)

	msg := "[INFO] Starting server on 8080."
	stat, _ := file.Stat()
	log.Println(msg)
	fmt.Println(msg, "Server log is available in", stat.Name())


	http.ListenAndServe(":8080", withLogging)
}