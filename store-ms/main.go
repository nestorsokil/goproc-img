package main

import (
	"google.golang.org/grpc"
	"github.com/nestorsokil/goproc-img/store-ms/api"
	"net"
	"log"
	"fmt"
	"os"
	"encoding/json"
	"path/filepath"
)

const (
	CONFIG_FILE = "conf.json"
	FS_SEP      = string(filepath.Separator)
)



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

	grpcServer := grpc.NewServer(
		// options :
		grpc.MaxRecvMsgSize(5 << 30),
	)
	api.RegisterStoreServiceServer(grpcServer, &storeServer{config:conf})
	log.Printf("Starting grpc server on port %v\n", conf.Port)
	grpcServer.Serve(listener)
}
