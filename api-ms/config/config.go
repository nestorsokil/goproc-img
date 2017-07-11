package config

import (
	"os"
	"log"
	"encoding/json"
)

const CONF_FILE = "config/conf.json"

type ApiConfig struct {
	StoreMsUrl string `json:"store_ms_url"`
}

func LoadConfig() ApiConfig {
	var conf ApiConfig
	file, err := os.Open(CONF_FILE)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	err = json.NewDecoder(file).Decode(&conf)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}
	return conf
}