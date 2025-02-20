// internal/service/product_test.go
package service_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/00Chaotic/flip-tech-test/backend/internal/model"
	"github.com/00Chaotic/flip-tech-test/backend/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetProducts(t *testing.T) {
	t.Parallel()

	products := []*model.Product{
		{
			SKU:       "sku",
			Name:      "name",
			Price:     1,
			Inventory: 1,
		},
	}

	type fields struct {
		MockOperations func(r *MockProductRepository)
	}

	type want struct {
		Error    string
		Products []*model.Product
		Status   int
	}

	testTable := map[string]struct {
		Fields fields
		Want   want
	}{
		"error - database error": {
			Fields: fields{func(r *MockProductRepository) {
				r.On("GetProducts", mock.AnythingOfType("context.backgroundCtx")).Return(([]*model.Product)(nil), errors.New("database error"))
			}},
			Want: want{
				Error:  "failed to get products",
				Status: http.StatusInternalServerError,
			},
		},
		"success": {
			Fields: fields{func(r *MockProductRepository) {
				r.On("GetProducts", mock.AnythingOfType("context.backgroundCtx")).Return(products, nil)
			}},
			Want: want{
				Products: products,
				Status:   http.StatusOK,
			},
		},
	}

	for name, tt := range testTable {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			mockRepository := &MockProductRepository{}
			productService := service.NewProductService(mockRepository)

			req, err := http.NewRequest(http.MethodGet, "/products", nil)
			assert.NoError(t, err)

			recorder := httptest.NewRecorder()
			handler := http.HandlerFunc(productService.GetProducts)

			tt.Fields.MockOperations(mockRepository)

			handler.ServeHTTP(recorder, req)

			var response model.ProductsResponse

			err = json.Unmarshal(recorder.Body.Bytes(), &response)
			require.NoError(t, err, "failed to unmarshal response body")

			if tt.Want.Error != "" {
				assert.Equal(t, tt.Want.Error, response.Error)
			}

			assert.Equal(t, tt.Want.Products, response.Products)
			assert.Equal(t, tt.Want.Status, recorder.Code)

			mockRepository.AssertExpectations(t)
		})
	}
}

// MockProductRepository is a mock implementation of the ProductRepository interface.
type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) GetProducts(ctx context.Context) ([]*model.Product, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*model.Product), args.Error(1)
}

func (m *MockProductRepository) UpdateProductInventories(ctx context.Context, items []model.PurchaseItem) (float64, []*model.Product, error) {
	args := m.Called(ctx, items)
	return args.Get(0).(float64), args.Get(1).([]*model.Product), args.Error(2)
}
