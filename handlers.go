package main

import (
	"github.com/plally/random_fox_api/foxes"
	"github.com/plally/random_fox_api/storage"
	"log"
	"math/rand"
	"net/http"
	"path"
)

func NewFoxServer(s storage.Storage) FoxServer{
	foxList, err := s.List()
	if err != nil {
		log.Fatal(err)
	}

	return FoxServer{
		store:   s,
		foxList: foxList,
	}
}

type FoxServer struct {
	store storage.Storage
	foxList []string
}

func (f FoxServer) random(w http.ResponseWriter, r *http.Request) {
	id := f.foxList[rand.Intn(len(f.foxList))]
	format := path.Ext(r.URL.Path)[1:]
	data, err := foxes.RetrieveFox(f.store, id, format)
	if err == foxes.InvalidFormatErr {
		w.Write([]byte("retrieving images in that format is not supported "+format))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil {
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", getContentType(r.URL.Path))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func getContentType(p string) string {
	switch path.Ext(p) {
	case "jpg": return "image/jpeg"
	case "png": return "image/png"
	case "json": return "application/json"
	}
	return ""
}