package handlers

import (
	"net/http"
	"product-api/entity"
)

// swagger:route GET / products listProducts
// returns a list of products
// responses:
//	200: productsResponse
func (p *ProductHandler) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := entity.GetProducts()
	rw.Header().Add("Content-Type", "application/json")
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to marshall json", http.StatusInternalServerError)
	}

}
