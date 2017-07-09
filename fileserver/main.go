package main

import (
	"os"
	"encoding/json"
	"log"
	"net/http"
	"fmt"
)

const CONFIG_FILE = "conf.json"

type config struct {
	Port int `json:"port"`
	Directory string `json:"directory"`
}

func main() {
	configFile, err := os.Open(CONFIG_FILE)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	conf := config{}
	err = json.NewDecoder(configFile).Decode(&conf)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}

	http.Handle("/", http.FileServer(http.Dir(conf.Directory)))
	log.Printf("Serving static files on %d\n", conf.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%v", conf.Port), nil)
	log.Fatal(err)
}
