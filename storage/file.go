package storage

import (
	"github.com/google/uuid"
	"io/ioutil"
	"os"
	"path"
)

type FileStore struct {
	Path string
}

func NewFileStore(path string) FileStore {
	os.MkdirAll(path, os.ModePerm)
	return FileStore{Path: path}
}

func (f FileStore) Create(imageData []byte) (string, error) {
	id := uuid.New().String()
	file, err := os.Create(path.Join(f.Path, id))
	defer file.Close()
	if err != nil { return id, err }

	_, err = file.Write(imageData)
	return id, err
}

func (f FileStore) Retrieve(id string) ([]byte, error) {
	file, err := os.Open(path.Join(f.Path, id))
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(file)
}

func (f FileStore) List() ([]string, error) {
	files, err := ioutil.ReadDir(f.Path)
	if err != nil {
		return nil, err
	}

	ids := make([]string, len(files))
	for i, file := range files {
		if file.IsDir() { continue }
		ids[i] = file.Name()
	}
	return ids, nil
}