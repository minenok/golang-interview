package main

import (
	"github.com/gorilla/mux"
	"github.com/minenok/golang-interview/handlers"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/healthz", handlers.Healthz).Methods(http.MethodGet)
	r.HandleFunc("/products", handlers.Products).Methods(http.MethodGet)
	r.HandleFunc("/products", handlers.NewProduct).Methods(http.MethodPost)

	http.Handle("/", r)
	if err := http.ListenAndServe(":8090", nil); err != nil {
		panic(err)
	}
}
