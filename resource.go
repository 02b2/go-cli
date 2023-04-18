package main

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
)

type Resource interface {
	Name() string
	Content() []byte
}

type StaticResource struct {
	StaticName    string
	StaticContent []byte
}

func (r *StaticResource) Name() string {
	return r.StaticName
}

func (r *StaticResource) Content() []byte {
	return r.StaticContent
}

func NewStaticResource(name string, content []byte) *StaticResource {
	return &StaticResource{
		StaticName:    name,
		StaticContent: content,
	}
}

func LoadResourceFromPath(path string) (Resource, error) {
	bytes, err := ioutil.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, err
	}
	name := filepath.Base(path)
	return NewStaticResource(name, bytes), nil
}

func LoadResourceFromURLString(urlStr string) (Resource, error) {
	res, err := http.Get(urlStr)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	name := filepath.Base(urlStr)
	return NewStaticResource(name, bytes), nil
}
