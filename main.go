package main

import (
	"flag"
	"github.com/plally/random_fox_api/storage"
	"net/http"
)

var (
	port = flag.String("addr", ":8080", "the address to listen on")
)

func main() {
	s := storage.NewFileStore("data/fox")
	f := NewFoxServer(s)
	flag.Parse()

	http.HandleFunc("/random.png", f.random)
	http.HandleFunc("/random.jpg", f.random)

	http.ListenAndServe(*port, nil)

}
