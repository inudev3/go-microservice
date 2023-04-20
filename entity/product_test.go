package entity

import "testing"

func TestValidation(t *testing.T) {
	p := &Product{Name: "nics", Price: 30, SKU: "def-ghi"}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
