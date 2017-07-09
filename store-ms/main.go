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
	"github.com/nu7hatch/gouuid"
)

const (
	CONFIG_FILE = "conf.json"
	FS_SEP      = string(filepath.Separator)
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
	var filename string
	if req.Filename == "" {
		id, err := uuid.NewV4()
		if err != nil {
			return nil, err
		}
		filename = id.String()
	} else {
		filename = req.Filename
	}

	today := time.Now().Local().Format("2006-01-02")
	out, err := s.getFile(today, filename)
	if err != nil {
		return notOk(), err
	}
	defer out.Close()

	_, err = out.Write(req.Data)
	if err != nil {
		return notOk(), err
	}

	relPath := today + "/" + filename
	url := s.config.StorageUrl + "/" + relPath
	return ok(url), nil
}

func (s *storeServer) getFile(folder, filename string) (*os.File, error) {
	dir := s.config.DataFolderPath + FS_SEP + folder
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0777)
	}

	path := dir + FS_SEP + filename
	out, err := os.OpenFile(path, os.O_CREATE, 0777)
	if err != nil {
		return nil, err
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
