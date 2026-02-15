package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"web-server/internal/database"
	"web-server/internal/models"

	"github.com/shopspring/decimal"
)

func TestCreateProduct(t *testing.T) {
	tests := []struct {
		name           string
		body           string
		mockCreate     func(ctx context.Context, product *models.Product) error
		expectedStatus int
		checkResponse  func(t *testing.T, body []byte)
	}{
		{
			name: "success - create product",
			body: `{"name":"Test Product","price":29.99,"stock":100}`,
			mockCreate: func(ctx context.Context, product *models.Product) error {
				product.ID = 1
				return nil
			},
			expectedStatus: http.StatusCreated,
			checkResponse: func(t *testing.T, body []byte) {
				var p models.Product
				if err := json.Unmarshal(body, &p); err != nil {
					t.Fatalf("failed to unmarshal response: %v", err)
				}
				if p.ID != 1 {
					t.Errorf("expected ID 1, got %d", p.ID)
				}
				if p.Name != "Test Product" {
					t.Errorf("expected name 'Test Product', got %s", p.Name)
				}
			},
		},
		{
			name:           "error - invalid json",
			body:           `{"name":"Test Product","price":invalid}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "error - negative price",
			body:           `{"name":"Test Product","price":-10.00,"stock":100}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "error - negative stock",
			body:           `{"name":"Test Product","price":10.00,"stock":-5}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "error - database failure",
			body: `{"name":"Test Product","price":29.99,"stock":100}`,
			mockCreate: func(ctx context.Context, product *models.Product) error {
				return errors.New("database connection failed")
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB := &database.MockService{
				CreateProductFunc: tt.mockCreate,
			}
			handler := NewProductHandler(mockDB)

			req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			handler.CreateProduct(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.checkResponse != nil {
				tt.checkResponse(t, rr.Body.Bytes())
			}
		})
	}
}

func TestGetProduct(t *testing.T) {
	tests := []struct {
		name           string
		id             string
		mockGet        func(ctx context.Context, id int) (*models.Product, error)
		expectedStatus int
		checkResponse  func(t *testing.T, body []byte)
	}{
		{
			name: "success - get existing product",
			id:   "1",
			mockGet: func(ctx context.Context, id int) (*models.Product, error) {
				return &models.Product{
					ID:       id,
					Name:     "Test Product",
					Price:    decimal.NewFromFloat(29.99),
					Stock:    100,
					IsActive: true,
				}, nil
			},
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, body []byte) {
				var p models.Product
				if err := json.Unmarshal(body, &p); err != nil {
					t.Fatalf("failed to unmarshal: %v", err)
				}
				if p.Name != "Test Product" {
					t.Errorf("expected 'Test Product', got %s", p.Name)
				}
			},
		},
		{
			name:           "error - invalid id",
			id:             "abc",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "error - product not found",
			id:   "999",
			mockGet: func(ctx context.Context, id int) (*models.Product, error) {
				return nil, errors.New("product not found")
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB := &database.MockService{
				GetProductFunc: tt.mockGet,
			}
			handler := NewProductHandler(mockDB)

			// Crear request con PathValue
			req := httptest.NewRequest(http.MethodGet, "/products/"+tt.id, nil)
			req.SetPathValue("id", tt.id)
			rr := httptest.NewRecorder()

			handler.GetProduct(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.checkResponse != nil {
				tt.checkResponse(t, rr.Body.Bytes())
			}
		})
	}
}

func TestGetAllProducts(t *testing.T) {
	tests := []struct {
		name           string
		mockGetAll     func(ctx context.Context) ([]models.Product, error)
		expectedStatus int
		expectedCount  int
	}{
		{
			name: "success - get all products",
			mockGetAll: func(ctx context.Context) ([]models.Product, error) {
				return []models.Product{
					{ID: 1, Name: "Product 1", Price: decimal.NewFromFloat(10.00)},
					{ID: 2, Name: "Product 2", Price: decimal.NewFromFloat(20.00)},
				}, nil
			},
			expectedStatus: http.StatusOK,
			expectedCount:  2,
		},
		{
			name: "success - empty list",
			mockGetAll: func(ctx context.Context) ([]models.Product, error) {
				return []models.Product{}, nil
			},
			expectedStatus: http.StatusOK,
			expectedCount:  0,
		},
		{
			name: "error - database failure",
			mockGetAll: func(ctx context.Context) ([]models.Product, error) {
				return nil, errors.New("database error")
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB := &database.MockService{
				GetAllProductsFunc: tt.mockGetAll,
			}
			handler := NewProductHandler(mockDB)

			req := httptest.NewRequest(http.MethodGet, "/products", nil)
			rr := httptest.NewRecorder()

			handler.GetAllProducts(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedStatus == http.StatusOK {
				var products []models.Product
				if err := json.Unmarshal(rr.Body.Bytes(), &products); err != nil {
					t.Fatalf("failed to unmarshal: %v", err)
				}
				if len(products) != tt.expectedCount {
					t.Errorf("expected %d products, got %d", tt.expectedCount, len(products))
				}
			}
		})
	}
}

func TestUpdateProduct(t *testing.T) {
	tests := []struct {
		name           string
		id             string
		body           string
		mockUpdate     func(ctx context.Context, product *models.Product) error
		expectedStatus int
	}{
		{
			name: "success - update product",
			id:   "1",
			body: `{"name":"Updated Product","price":39.99,"stock":50}`,
			mockUpdate: func(ctx context.Context, product *models.Product) error {
				if product.ID != 1 {
					return errors.New("wrong ID")
				}
				return nil
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "error - invalid id",
			id:             "abc",
			body:           `{"name":"Test","price":10.00}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "error - invalid json",
			id:             "1",
			body:           `{"name":"Test","price":invalid}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "error - negative price",
			id:             "1",
			body:           `{"name":"Test","price":-10.00}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "error - database failure",
			id:   "1",
			body: `{"name":"Test","price":10.00}`,
			mockUpdate: func(ctx context.Context, product *models.Product) error {
				return errors.New("update failed")
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB := &database.MockService{
				UpdateProductFunc: tt.mockUpdate,
			}
			handler := NewProductHandler(mockDB)

			req := httptest.NewRequest(http.MethodPut, "/products/"+tt.id, bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")
			req.SetPathValue("id", tt.id)
			rr := httptest.NewRecorder()

			handler.UpdateProduct(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rr.Code)
			}
		})
	}
}

func TestDeleteProduct(t *testing.T) {
	tests := []struct {
		name           string
		id             string
		mockDelete     func(ctx context.Context, id int) error
		expectedStatus int
	}{
		{
			name: "success - delete product",
			id:   "1",
			mockDelete: func(ctx context.Context, id int) error {
				return nil
			},
			expectedStatus: http.StatusNoContent,
		},
		{
			name:           "error - invalid id",
			id:             "abc",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "error - database failure",
			id:   "1",
			mockDelete: func(ctx context.Context, id int) error {
				return errors.New("delete failed")
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB := &database.MockService{
				DeleteProductFunc: tt.mockDelete,
			}
			handler := NewProductHandler(mockDB)

			req := httptest.NewRequest(http.MethodDelete, "/products/"+tt.id, nil)
			req.SetPathValue("id", tt.id)
			rr := httptest.NewRecorder()

			handler.DeleteProduct(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rr.Code)
			}
		})
	}
}
