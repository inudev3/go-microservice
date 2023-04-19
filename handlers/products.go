package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"microservice/entity"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *ProductHandler {
	return &ProductHandler{l}
}

func (p *ProductHandler) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := entity.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to marshall json", http.StatusInternalServerError)
	}

}
func (p *ProductHandler) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("handle POST product")
	prod := &entity.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "unable to unmarshall json", http.StatusBadRequest)
		return
	}
	entity.AddProduct(prod)
	lp := entity.GetProducts()
	err = lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to marshall json", http.StatusInternalServerError)
	}
}

func (p *ProductHandler) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to unmarshall json", http.StatusBadRequest)
		return
	}
	p.l.Println("Handle PUT Update", id)
	prod := &entity.Product{}
	err = prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshall json", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(rw, "Invalid ID", http.StatusBadRequest)
		return
	}
	err = entity.UpdateProduct(id, prod)
	if err == entity.ErrorProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct{}

func (p *ProductHandler) MiddlewareValidation(next http.Handler) http.Handler {
	var myfunc http.HandlerFunc

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := entity.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}
		ctx := r.Context().Value(KeyProduct{})

	})
}
