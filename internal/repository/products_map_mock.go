package repository

import "app/internal"

func NewProductsMapMock() *ProductsMapMock {
	return &ProductsMapMock{}
}

type ProductsMapMock struct {
	SearchProductsFunc func(query internal.ProductQuery) (map[int]internal.Product, error)
}

func (m *ProductsMapMock) SearchProducts(query internal.ProductQuery) (map[int]internal.Product, error) {
	return m.SearchProductsFunc(query)
}
