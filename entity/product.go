package entity

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"SKU"`
	CreatedAt   string  `json:"-"`
	UpdatedAt   string  `json:"-"`
	DeletedAt   string  `json:"-"`
}
type Products []*Product

func (p Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

var productList = Products{

	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffe",
		Price:       2.45,
		SKU:         "abc323",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "bitter taste",
		Price:       1.95,
		SKU:         "fbcj124",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
}

func GetProducts() Products {
	return productList
}
func AddProduct(prod *Product) {
	prod.ID = NextID()
	productList = append(productList, prod)
}
func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[pos] = p
	return nil
}
func NextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

var ErrorProductNotFound = fmt.Errorf("Product Not Found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrorProductNotFound
}
