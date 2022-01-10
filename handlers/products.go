package handlers

import (
	"fmt"
	"log"
)

// Products is a http.Handler
type Products struct {
	l *log.Logger
}

// func (p *Products) ListAll(rw http.ResponseWriter, r *http.Request) {
// 	p.l.Println("Handle GET Products")

// 	// fetch the products from the datastore
// 	lp := data.GetProducts()

// 	// serialize the list to JSON
// 	err := lp.ToJSON(rw)
// 	if err != nil {
// 		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
// 	}
// }

// NewProducts returns a new products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// KeyProduct is product key for add new product
type KeyProduct struct{}

// ErrInvalidProductPath is an error message when the product path is not valid
var ErrInvalidProductPath = fmt.Errorf("Invalid Path, path should be /products/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}
