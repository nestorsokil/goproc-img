package api

import (
	"google.golang.org/grpc"
	"golang.org/x/net/context"
)

type StoreClient struct {
	connection *grpc.ClientConn
	client StoreServiceClient
}

func (c StoreClient) SaveFile(in *SaveRequest)(*SaveResult, error) {
	return c.client.SaveFile(context.Background(), in)
}

func (c StoreClient) CloseClient() error {
	return c.connection.Close()
}

func NewClient(serverUrl string) (*StoreClient, error) {
	conn, err := grpc.Dial(serverUrl, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	clientImpl := NewStoreServiceClient(conn)
	return &StoreClient{conn, clientImpl}, nil
}
