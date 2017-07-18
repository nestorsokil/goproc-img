package config

import (
	"os"
	"log"
	"encoding/json"
	"io/ioutil"
)

const CONF_FILE = "config/conf.json"

type ApiConfig struct {
	StoreMsUrl string `json:"store_ms_url"`
	LogFilePath string `json:"log_file_path"`
	EnableJWTSecurity bool `json:"enable_jwt_security"`
	Audience []string `json:"audience"`
	PrivateKeyPath string `json:"private_key_path"`
	PublicKeyPath string `json:"public_key_path"`
	Auth0DomainName string `json:"auth_0_domain_name"`
	Port string `json:"port"`
}

var Settings = LoadConfig()
var PrivateKey []byte = load(Settings.PrivateKeyPath)
var PublicKey  []byte = load(Settings.PublicKeyPath)

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

func load(path string) []byte {
	if !Settings.EnableJWTSecurity {
		return nil
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return data
}