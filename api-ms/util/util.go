package util

import (
	"github.com/nestorsokil/goproc-img/store-ms/api"
	"net/http"
	"io"
	"google.golang.org/grpc"
	"context"
	"time"
	"encoding/json"
)

func GetClient(url string) (api.StoreServiceClient, error) {
	ctx, onCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer onCancel()
	conn, err := grpc.DialContext(ctx, url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return api.NewStoreServiceClient(conn), nil
}

func Respond400(writer http.ResponseWriter, msg string) {
	writer.WriteHeader(http.StatusBadRequest)
	writer.Header().Set("Content-Type", "text/plain")
	io.WriteString(writer, msg)
}

func Respond403(writer http.ResponseWriter, msg string) {
	writer.WriteHeader(http.StatusForbidden)
	writer.Header().Set("Content-Type", "text/plain")
	io.WriteString(writer, msg)
}

func Respond500(writer http.ResponseWriter, msg string) {
	writer.WriteHeader(http.StatusInternalServerError)
	writer.Header().Set("Content-Type", "text/plain")
	io.WriteString(writer, msg)
}

func RespondJson(writer http.ResponseWriter, data interface{}) {
	marshaled, err :=  json.Marshal(data)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(marshaled)
}