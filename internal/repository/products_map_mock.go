package repository

import "app/internal"

func NewProductsMapMock() *ProductsMapMock {
	return &ProductsMapMock{}
}

type ProductsMapMock struct {
	SearchProductsFunc func(query internal.ProductQuery) (map[int]internal.Product, error)
	Spy                struct {
		SearchProductsCallCount int
		SearchProductsArgs      []internal.ProductQuery
	}
}

func (m *ProductsMapMock) SearchProducts(query internal.ProductQuery) (map[int]internal.Product, error) {
	m.Spy.SearchProductsCallCount++
	m.Spy.SearchProductsArgs = append(m.Spy.SearchProductsArgs, query)

	return m.SearchProductsFunc(query)
}
