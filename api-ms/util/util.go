package util

import (
	"github.com/nestorsokil/goproc-img/store-ms/api"
	"net/http"
	"io"
	"google.golang.org/grpc"
)

func GetClient(url string) (api.StoreServiceClient, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
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
