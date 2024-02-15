package internal

import "errors"

// ProductAttributes is an struct that represents a product.
type ProductAttributes struct {
	// Description is the description of the product.
	Description string
	// Price is the price of the product.
	Price float64
	// SellerId is the id of the seller of the product.
	SellerId int
}

// Product is an struct that represents a product in the storage.
type Product struct {
	// Id is the unique identifier of the product.
	Id int
	// ProductAttributes is the attributes of the product.
	ProductAttributes
}

var (
	// ErrProductNotFound is the error returned when the product is not found.
	ErrProductNotFound = errors.New("product not found")
	ErrInvalidQuery    = errors.New("invalid query")
)

// ProductQuery is an struct that represents a query to the storage.
type ProductQuery struct {
	// Id is the unique identifier of the product.
	Id int
}
