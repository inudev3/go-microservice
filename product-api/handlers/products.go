//	Package classification Product API
//
//	Documentation for Product API
// 	Schemes: http
// 	BasePath: /
// 	Version: 1.0.0
// 	License: MIT http://opensource.org/licenses/MIT
// 	Contact: John Doe <john.doe@example.com>
//
// 	Consumes:
// 	- application/json
//
// 	Produces:
// 	- application/json
//
//	swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"product-api/entity"
)

type ProductHandler struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *ProductHandler {
	return &ProductHandler{l}
}


type KeyProduct struct{}

func (p *ProductHandler) MiddlewareValidation(next http.Handler) http.Handler {

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &entity.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}
		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating product", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating product: %s", err),
				http.StatusBadRequest,
			)
			return
		}
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)
		next.ServeHTTP(rw, req)
	})
}
