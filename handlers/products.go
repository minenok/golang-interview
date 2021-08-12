package handlers

import (
	"encoding/json"
	"github.com/minenok/golang-interview/db"
	"net/http"
)

func NewProduct(w http.ResponseWriter, r *http.Request) {
	dbc := db.NewDB()
	if err := dbc.SaveProduct(r.FormValue("name"), r.FormValue("description")); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

type outProduct struct {
	id          int    `json:"int"`
	name        string `json:"name"`
	description string `json:"description"`
}

func Products(w http.ResponseWriter, r *http.Request) {
	dbc := db.NewDB()
	rows, err := dbc.Products()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	products := make([]outProduct, 0)
	for rows.Next() {
		p := outProduct{}
		_ = rows.Scan(&p.id, &p.name, &p.description)
		products = append(products, p)
	}

	w.WriteHeader(http.StatusOK)
	bs, _ := json.Marshal(products)
	w.Write(bs)
}
