package handler_test

import (
	"app/internal"
	"app/internal/handler"
	"app/internal/repository"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGet_NoIdProvided(t *testing.T) {
	t.Run("no id provided", func(t *testing.T) {
		rp := repository.NewProductsMapMock()
		hd := handler.NewProductsDefault(rp)

		hdFunc := hd.Get()

		req := httptest.NewRequest(http.MethodGet, "/product", nil)
		res := httptest.NewRecorder()

		hdFunc(res, req)

		require.Equal(t, http.StatusBadRequest, res.Code)
		require.Equal(t, `{"status":"Bad Request","message":"no id provided"}`, res.Body.String())
	})
}

func TestGet(t *testing.T) {
	testCases := []struct {
		name        string
		id          string
		queryReturn map[int]internal.Product
		queryErr    error
		expectCode  int
		expectBody  string
	}{
		{
			name:       "invalid id",
			id:         "test",
			expectCode: http.StatusBadRequest,
			expectBody: `{"status":"Bad Request","message":"invalid id"}`,
		}, {
			name:       "invalid query",
			id:         "0",
			queryErr:   internal.ErrInvalidQuery,
			expectCode: http.StatusBadRequest,
			expectBody: `{"status":"Bad Request","message":"invalid query"}`,
		}, {
			name:       "product not found",
			id:         "100",
			queryErr:   internal.ErrProductNotFound,
			expectCode: http.StatusNotFound,
			expectBody: `{"status":"Not Found","message":"product not found"}`,
		}, {
			name:       "internal error",
			id:         "1",
			queryErr:   errors.New("internal error"),
			expectCode: http.StatusInternalServerError,
			expectBody: `{"status":"Internal Server Error","message":"internal error"}`,
		}, {
			name: "success",
			id:   "1",
			queryReturn: map[int]internal.Product{
				1: {
					Id: 1,
					ProductAttributes: internal.ProductAttributes{
						Description: "test",
						Price:       100,
						SellerId:    1,
					},
				},
			},
			expectCode: http.StatusOK,
			expectBody: `{
				"message": "success",
				"data": {
					"1": {
						"id": 1,
						"description": "test",
						"price": 100,
						"seller_id": 1
						}
					}
				}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rp := repository.NewProductsMapMock()
			searchFunc := func(query internal.ProductQuery) (map[int]internal.Product, error) {
				return tc.queryReturn, tc.queryErr
			}
			rp.SearchProductsFunc = searchFunc
			hd := handler.NewProductsDefault(rp)

			hdFunc := hd.Get()

			req := httptest.NewRequest(http.MethodGet, "/product", nil)
			query := req.URL.Query()
			query.Add("id", tc.id)
			req.URL.RawQuery = query.Encode()

			res := httptest.NewRecorder()

			hdFunc(res, req)

			require.Equal(t, tc.expectCode, res.Code)
			require.JSONEq(t, tc.expectBody, res.Body.String())
		})
	}
}
