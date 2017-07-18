package main

import (

	"github.com/nestorsokil/goproc-img/api-ms/handlers"
	"log"
	"net/http"
	"os"
	"fmt"
	"github.com/gorilla/mux"
	wrappers "github.com/gorilla/handlers"
	"github.com/nestorsokil/goproc-img/api-ms/config"
	"path/filepath"
)

const (
	PING_PATH   = "/ping"
	LOGIN_PATH  = "/login"
	UPLOAD_PATH = "/api/v1/upload"
)

func main() {
	logDir := filepath.Dir(config.Settings.LogFilePath)
	if logDir != "" {
		err := os.MkdirAll(logDir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
	logFile, err := os.OpenFile(config.Settings.LogFilePath, os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()


	log.SetOutput(logFile)
	router := mux.NewRouter()

	router.Handle(PING_PATH, handlers.DoPong()).Methods("GET")
	router.Handle(UPLOAD_PATH, handlers.StoreFileByUrl()).Methods("GET")
	router.Handle(UPLOAD_PATH, handlers.StoreFileByPostData()).Methods("POST")
	router.Handle(LOGIN_PATH, handlers.DoLogin()).Methods("POST")

	withLogging := wrappers.LoggingHandler(logFile, router)

	msg := fmt.Sprintf("[INFO] Starting server on %v.", config.Settings.Port)
	log.Println(msg)
	fmt.Println(msg, "Server log is available in", config.Settings.LogFilePath)


	http.ListenAndServe(config.Settings.Port, withLogging)
}