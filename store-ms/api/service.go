package api

import "golang.org/x/net/context"

type storeService struct {

}

func (s *storeService) SaveFile(ctx context.Context, req *SaveRequest) (*SaveResult, error) {
	return &SaveResult{true, nil}, nil
}