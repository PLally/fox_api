package main

import (
	"flag"
	"github.com/plally/random_fox_api/storage"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	port = flag.String("addr", ":8080", "the address to listen on")
	dir = flag.String("dir", "data/fox",  "directory for fox")
)

func main() {
	flag.Parse()
	r := mux.NewRouter()
	s := storage.NewFileStore(*dir)
	f := NewFoxServer(s)

	r.HandleFunc("/random.{format:[a-z]+}", f.random).Methods("GET", "OPTIONS")
	r.HandleFunc("/{id}.{format:[a-z]+}", f.get).Methods("GET", "OPTIONS")

	r.Use(mux.CORSMethodMiddleware(r), SetHeader("Access-Control-Allow-Origin", "*"))
	http.ListenAndServe(*port, r)

}

func SetHeader(key, value string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(key, value)
			next.ServeHTTP(w, r)
		})
	}
}
