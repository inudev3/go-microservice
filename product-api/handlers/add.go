package handlers

import (
	"net/http"
	"product-api/entity"
)
// swagger:route POST /product products addProduct
// responses:
//  200: productResponse
func (p *ProductHandler) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("handle POST product")
	prod := r.Context().Value(KeyProduct{}).(entity.Product)
	entity.AddProduct(&prod)
}

