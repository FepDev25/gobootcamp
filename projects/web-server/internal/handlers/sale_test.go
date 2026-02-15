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

func intPtr(i int) *int {
	return &i
}

func TestCreateSale(t *testing.T) {
	tests := []struct {
		name           string
		body           string
		mockCreate     func(ctx context.Context, req *models.CreateSaleRequest) (*models.Sale, error)
		expectedStatus int
		checkResponse  func(t *testing.T, body []byte)
	}{
		{
			name: "success - create sale with single item",
			body: `{"customer_id":1,"payment_method":"CASH","items":[{"product_id":1,"quantity":2}]}`,
			mockCreate: func(ctx context.Context, req *models.CreateSaleRequest) (*models.Sale, error) {
				return &models.Sale{
					ID:            1,
					CustomerID:    req.CustomerID,
					PaymentMethod: req.PaymentMethod,
					TotalAmount:   decimal.NewFromFloat(59.98),
					Items: []models.SaleItem{
						{ID: 1, ProductID: 1, Quantity: 2, UnitPrice: decimal.NewFromFloat(29.99)},
					},
				}, nil
			},
			expectedStatus: http.StatusCreated,
			checkResponse: func(t *testing.T, body []byte) {
				var s models.Sale
				if err := json.Unmarshal(body, &s); err != nil {
					t.Fatalf("failed to unmarshal: %v", err)
				}
				if s.ID != 1 {
					t.Errorf("expected sale ID 1, got %d", s.ID)
				}
				if s.PaymentMethod != "CASH" {
					t.Errorf("expected payment method CASH, got %s", s.PaymentMethod)
				}
			},
		},
		{
			name: "success - create sale with multiple items",
			body: `{"customer_id":2,"payment_method":"CARD","items":[{"product_id":1,"quantity":1},{"product_id":2,"quantity":3}]}`,
			mockCreate: func(ctx context.Context, req *models.CreateSaleRequest) (*models.Sale, error) {
				return &models.Sale{
					ID:            2,
					CustomerID:    req.CustomerID,
					PaymentMethod: req.PaymentMethod,
					TotalAmount:   decimal.NewFromFloat(149.95),
				}, nil
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name:           "error - invalid json",
			body:           `{"customer_id":1,"items":invalid}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "error - empty items",
			body:           `{"customer_id":1,"payment_method":"CASH","items":[]}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "error - insufficient stock",
			body: `{"customer_id":1,"payment_method":"CASH","items":[{"product_id":1,"quantity":100}]}`,
			mockCreate: func(ctx context.Context, req *models.CreateSaleRequest) (*models.Sale, error) {
				return nil, errors.New("insufficient stock for product 1")
			},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name: "error - product not found",
			body: `{"customer_id":1,"payment_method":"CASH","items":[{"product_id":999,"quantity":1}]}`,
			mockCreate: func(ctx context.Context, req *models.CreateSaleRequest) (*models.Sale, error) {
				return nil, errors.New("product 999 not found")
			},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name: "error - database transaction failure",
			body: `{"customer_id":1,"payment_method":"CASH","items":[{"product_id":1,"quantity":1}]}`,
			mockCreate: func(ctx context.Context, req *models.CreateSaleRequest) (*models.Sale, error) {
				return nil, errors.New("transaction failed: connection lost")
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB := &database.MockService{
				CreateSaleFunc: tt.mockCreate,
			}
			handler := NewSaleHandler(mockDB)

			req := httptest.NewRequest(http.MethodPost, "/sales", bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			handler.CreateSale(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.checkResponse != nil {
				tt.checkResponse(t, rr.Body.Bytes())
			}
		})
	}
}

func TestGetSale(t *testing.T) {
	tests := []struct {
		name           string
		id             string
		mockGet        func(ctx context.Context, id int) (*models.Sale, error)
		expectedStatus int
		checkResponse  func(t *testing.T, body []byte)
	}{
		{
			name: "success - get sale with items",
			id:   "1",
			mockGet: func(ctx context.Context, id int) (*models.Sale, error) {
				return &models.Sale{
					ID:            id,
					CustomerID:    intPtr(1),
					PaymentMethod: "CASH",
					TotalAmount:   decimal.NewFromFloat(99.99),
					Items: []models.SaleItem{
						{
							ID:        1,
							SaleID:    id,
							ProductID: 1,
							Quantity:  2,
							UnitPrice: decimal.NewFromFloat(49.995),
							Subtotal:  decimal.NewFromFloat(99.99),
						},
					},
				}, nil
			},
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, body []byte) {
				var s models.Sale
				if err := json.Unmarshal(body, &s); err != nil {
					t.Fatalf("failed to unmarshal: %v", err)
				}
				if len(s.Items) != 1 {
					t.Errorf("expected 1 item, got %d", len(s.Items))
				}
				if s.Items[0].Quantity != 2 {
					t.Errorf("expected quantity 2, got %d", s.Items[0].Quantity)
				}
			},
		},
		{
			name:           "error - invalid id",
			id:             "abc",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "error - sale not found",
			id:   "999",
			mockGet: func(ctx context.Context, id int) (*models.Sale, error) {
				return nil, errors.New("sale not found")
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB := &database.MockService{
				GetSaleFunc: tt.mockGet,
			}
			handler := NewSaleHandler(mockDB)

			req := httptest.NewRequest(http.MethodGet, "/sales/"+tt.id, nil)
			req.SetPathValue("id", tt.id)
			rr := httptest.NewRecorder()

			handler.GetSale(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.checkResponse != nil {
				tt.checkResponse(t, rr.Body.Bytes())
			}
		})
	}
}

func TestGetSalesReport(t *testing.T) {
	tests := []struct {
		name           string
		mockGetReport  func(ctx context.Context) ([]models.SalesReportItem, error)
		expectedStatus int
		expectedCount  int
	}{
		{
			name: "success - get sales report",
			mockGetReport: func(ctx context.Context) ([]models.SalesReportItem, error) {
				return []models.SalesReportItem{
					{Date: "2024-01-15", TotalSales: decimal.NewFromFloat(1250.50)},
					{Date: "2024-01-14", TotalSales: decimal.NewFromFloat(980.00)},
					{Date: "2024-01-13", TotalSales: decimal.NewFromFloat(1500.75)},
				}, nil
			},
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name: "success - empty report",
			mockGetReport: func(ctx context.Context) ([]models.SalesReportItem, error) {
				return []models.SalesReportItem{}, nil
			},
			expectedStatus: http.StatusOK,
			expectedCount:  0,
		},
		{
			name: "error - database failure",
			mockGetReport: func(ctx context.Context) ([]models.SalesReportItem, error) {
				return nil, errors.New("failed to query sales data")
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB := &database.MockService{
				GetSalesReportFunc: tt.mockGetReport,
			}
			handler := NewSaleHandler(mockDB)

			req := httptest.NewRequest(http.MethodGet, "/sales/report", nil)
			rr := httptest.NewRecorder()

			handler.GetSalesReport(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedStatus == http.StatusOK {
				contentType := rr.Header().Get("Content-Type")
				if contentType != "application/json" {
					t.Errorf("expected Content-Type application/json, got %s", contentType)
				}

				var report []models.SalesReportItem
				if err := json.Unmarshal(rr.Body.Bytes(), &report); err != nil {
					t.Fatalf("failed to unmarshal: %v", err)
				}
				if len(report) != tt.expectedCount {
					t.Errorf("expected %d items, got %d", tt.expectedCount, len(report))
				}
			}
		})
	}
}

func TestSaleRequestValidation(t *testing.T) {
	tests := []struct {
		name           string
		body           string
		expectedStatus int
	}{
		{
			name:           "missing items",
			body:           `{"customer_id":1,"payment_method":"CASH"}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "empty items array",
			body:           `{"customer_id":1,"payment_method":"CASH","items":[]}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "invalid json",
			body:           `{"customer_id":1,"payment_method":"CASH","items":invalid}`,
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB := &database.MockService{
				CreateSaleFunc: func(ctx context.Context, req *models.CreateSaleRequest) (*models.Sale, error) {
					// Esta función no debería ser llamada si la validación funciona
					t.Error("CreateSale should not be called for invalid request")
					return nil, errors.New("should not reach here")
				},
			}
			handler := NewSaleHandler(mockDB)

			req := httptest.NewRequest(http.MethodPost, "/sales", bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			handler.CreateSale(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rr.Code)
			}
		})
	}
}
