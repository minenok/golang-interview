package cmd

import (
	"github.com/minenok/golang-interview/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/healthz", handlers.Healthz).Methods(http.MethodGet)
	r.HandleFunc("/products", handlers.Products).Methods(http.MethodGet)
	r.HandleFunc("/products", handlers.NewProduct).Methods(http.MethodPost)

	if err := http.ListenAndServe(":8090", nil); err != nil {
		panic(err)
	}
}
