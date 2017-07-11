package main

import (
	"github.com/nestorsokil/goproc-img/store-ms/api"
	"github.com/nu7hatch/gouuid"
	"time"
	"os"
	"golang.org/x/net/context"
)

type config struct {
	Port int `json:"port"`
	StorageUrl string `json:"storage_url"`
	DataFolderPath string `json:"data_folder_path"`
}

type storeServer struct {
	config config
}

func (s *storeServer) SaveFile(ctx context.Context, req *api.SaveRequest) (*api.SaveResult, error) {
	filename, err := getFilename(req)
	if err != nil {
		return nil, err
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
	url := s.config.StorageUrl + "/" + today + "/" + filename
	return ok(url), nil
}

func getFilename(request *api.SaveRequest) (string, error) {
	if request.Filename != "" {
		return request.Filename, nil
	}
	uid, err := createRandomFileName()
	if err != nil {
		return "", err
	}
	return uid, nil
}

func createRandomFileName() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	return id.String(), nil
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
