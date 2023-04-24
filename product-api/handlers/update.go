package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"product-api/entity"
	"strconv"
)

// swagger:route PUT /product/{id} products updateProduct
// update a product
// responses:
//	201: noContent
func (p *ProductHandler) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Invalid ID", http.StatusBadRequest)
		return
	}
	p.l.Println("Handle PUT Update", id)
	prod := r.Context().Value(KeyProduct{}).(entity.Product)

	err = entity.UpdateProduct(id, &prod)
	if err == entity.ErrorProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
