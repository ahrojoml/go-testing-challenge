package repository

import "app/internal"

// NewProductsMap returns a new ProductsMap.
func NewProductsMap(db map[int]internal.Product) *ProductsMap {
	// default values
	defaultDb := make(map[int]internal.Product)
	if db != nil {
		defaultDb = db
	}

	return &ProductsMap{
		db: defaultDb,
	}
}

// ProductAttributes is an struct that implements the RepositoryProducts interface.
type ProductsMap struct {
	// db is the map of products.
	db map[int]internal.Product
}

// SearchProducts returns a list of products that match the query.
func (r *ProductsMap) SearchProducts(query internal.ProductQuery) (map[int]internal.Product, error) {
	if query.Id <= 0 {
		return nil, internal.ErrInvalidQuery
	}

	p := make(map[int]internal.Product)

	// search the products
	for k, v := range r.db {
		// check if each query field is set
		if query.Id == v.Id {
			p[k] = v
			break
		}
	}

	if len(p) == 0 {
		return nil, internal.ErrProductNotFound
	}

	return p, nil
}
