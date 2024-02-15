package repository_test

import (
	"app/internal"
	"app/internal/repository"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearchProducts(t *testing.T) {
	testCases := []struct {
		name      string
		db        map[int]internal.Product
		query     internal.ProductQuery
		expectErr error
		expect    map[int]internal.Product
	}{
		{
			name:      "not found",
			db:        map[int]internal.Product{},
			query:     internal.ProductQuery{Id: 100},
			expectErr: internal.ErrProductNotFound,
		}, {
			name:      "null query",
			db:        map[int]internal.Product{},
			query:     internal.ProductQuery{Id: 0},
			expectErr: internal.ErrInvalidQuery,
		}, {
			name:   "success",
			db:     map[int]internal.Product{1: {Id: 1}},
			query:  internal.ProductQuery{Id: 1},
			expect: map[int]internal.Product{1: {Id: 1}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// setup
			rp := repository.NewProductsMap(tc.db)

			result, err := rp.SearchProducts(tc.query)

			if tc.expectErr != nil {
				require.ErrorIs(t, err, tc.expectErr)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expect, result)
			}
		})
	}
}
