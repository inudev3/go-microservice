package handlers

import (
	"product-api/entity"
)

// A list of products returns in the response
// swagger:response productsResponse
type productsResponse struct {
	// All products in the system
	// in: body
	Body []entity.Product
}
// swagger:response productResponse
type productResponse struct {
	// product that is added
	// in: body
	Body entity.Product
}

// swagger:response noContent
type noContent struct{}

// swagger:parameters updateProduct
type productIdParamWrapper struct {
	// the id of the product
	// in: path
	// required: true
	ID int `json:"id"`
}
