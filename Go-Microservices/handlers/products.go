package handlers

import (
	"context"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"microservices/data"
	"net/http"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJson(rw)
	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")
	prod := r.Context().Value(KeyProducts{}).(*data.Product)
	p.l.Printf("Product: %#v", prod)
	data.AddProducts(prod)
}

func (p *Products) Updateproducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "invalid id", http.StatusBadRequest)
	}
	p.l.Println("Handle PUT Products", id)
	prod := r.Context().Value(KeyProducts{}).(*data.Product)
	p.l.Printf("Product: %#v", prod)
	err = data.UpdateProducts(id, prod)
	if err != nil {
		http.Error(rw, "unable to update products", http.StatusInternalServerError)
	}
	if errors.Is(err, data.ErrProductNotFound) {
		http.Error(rw, "product not found", http.StatusNotFound)
	}
}

type KeyProducts struct{}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}
		err := prod.FromJson(r.Body)
		if err != nil {
			http.Error(w, "unable to unmarshal json", http.StatusBadRequest)
		}
		ctx := context.WithValue(r.Context(), KeyProducts{}, prod)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
