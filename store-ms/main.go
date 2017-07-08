package main

import (
	"google.golang.org/grpc"
	"github.com/nestorsokil/goproc-img/store-ms/api"
	"net"
	"log"
	"golang.org/x/net/context"
	"fmt"
	"os"
	"encoding/json"
	"time"
	"path/filepath"
)

const (
	CONFIG_FILE = "conf.json"
	SLASH = string(filepath.Separator)
)

type config struct {
	Port int `json:"port"`
	StorageUrl string `json:"storage_url"`
	DataFolderPath string `json:"data_folder_path"`
}

// TODO: for some reason server cannot be defined in a separate file
type storeServer struct {
	config config
}

func (s *storeServer) SaveFile(ctx context.Context, req *api.SaveRequest) (*api.SaveResult, error) {
	out, err := s.getFile(req.Filename)
	if err != nil {
		return notOk(), err
	}

	_, err = out.Write(req.Data)
	if err != nil {
		return notOk(), err
	}

	url := s.config.StorageUrl + "/" + s.config.DataFolderPath + "/" + req.Filename
	return ok(url), nil
}

func (s *storeServer) getFile(filename string) (*os.File, error) {
	today := time.Now().Local().Format("yyyy-MM-dd")
	folder := s.config.DataFolderPath + SLASH + today

	if _, err := os.Stat(folder); os.IsNotExist(err) {
		os.MkdirAll(folder, 0777)
	}

	path := folder + SLASH + filename
	var out *os.File
	defer out.Close()
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		out, err = os.Create(path)
		if err != nil {
			return nil, err
		}
	} else {
		out, err = os.Open(path)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

func ok(url string) *api.SaveResult {
	return &api.SaveResult{IsOk:true, URL:url}
}

func notOk() *api.SaveResult {
	return &api.SaveResult{IsOk:false, URL:""}
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
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Port))
	if err != nil {
		log.Fatalf("Failed starting a listener on port %v.\nError: %v", conf.Port, err)
	}
	grpcServer := grpc.NewServer()
	api.RegisterStoreServiceServer(grpcServer, &storeServer{config:conf})
	log.Printf("Starting grpc server on port %v\n", conf.Port)
	grpcServer.Serve(listener)
}
