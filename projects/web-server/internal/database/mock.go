package database

import (
	"context"
	"web-server/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shopspring/decimal"
)

// MockService implementa la interfaz Service para testing
type MockService struct {
	HealthFunc            func() map[string]string
	CloseFunc             func()
	GetPoolFunc           func() *pgxpool.Pool
	CreateCategoryFunc    func(ctx context.Context, category *models.Category) error
	GetCategoryFunc       func(ctx context.Context, id int) (*models.Category, error)
	GetAllCategoriesFunc  func(ctx context.Context) ([]models.Category, error)
	UpdateCategoryFunc    func(ctx context.Context, category *models.Category) error
	DeleteCategoryFunc    func(ctx context.Context, id int) error
	CreateProductFunc     func(ctx context.Context, product *models.Product) error
	GetProductFunc        func(ctx context.Context, id int) (*models.Product, error)
	GetAllProductsFunc    func(ctx context.Context) ([]models.Product, error)
	UpdateProductFunc     func(ctx context.Context, product *models.Product) error
	DeleteProductFunc     func(ctx context.Context, id int) error
	CreateCustomerFunc    func(ctx context.Context, customer *models.Customer) error
	GetCustomerFunc       func(ctx context.Context, id int) (*models.Customer, error)
	GetAllCustomersFunc   func(ctx context.Context) ([]models.Customer, error)
	UpdateCustomerFunc    func(ctx context.Context, customer *models.Customer) error
	DeleteCustomerFunc    func(ctx context.Context, id int) error
	CreateSaleFunc        func(ctx context.Context, req *models.CreateSaleRequest) (*models.Sale, error)
	GetSaleFunc           func(ctx context.Context, id int) (*models.Sale, error)
	GetSalesReportFunc    func(ctx context.Context) ([]models.SalesReportItem, error)
}

func (m *MockService) Health() map[string]string {
	if m.HealthFunc != nil {
		return m.HealthFunc()
	}
	return map[string]string{"status": "up"}
}

func (m *MockService) Close() {
	if m.CloseFunc != nil {
		m.CloseFunc()
	}
}

func (m *MockService) GetPool() *pgxpool.Pool {
	if m.GetPoolFunc != nil {
		return m.GetPoolFunc()
	}
	return nil
}

func (m *MockService) CreateCategory(ctx context.Context, category *models.Category) error {
	if m.CreateCategoryFunc != nil {
		return m.CreateCategoryFunc(ctx, category)
	}
	return nil
}

func (m *MockService) GetCategory(ctx context.Context, id int) (*models.Category, error) {
	if m.GetCategoryFunc != nil {
		return m.GetCategoryFunc(ctx, id)
	}
	return nil, nil
}

func (m *MockService) GetAllCategories(ctx context.Context) ([]models.Category, error) {
	if m.GetAllCategoriesFunc != nil {
		return m.GetAllCategoriesFunc(ctx)
	}
	return nil, nil
}

func (m *MockService) UpdateCategory(ctx context.Context, category *models.Category) error {
	if m.UpdateCategoryFunc != nil {
		return m.UpdateCategoryFunc(ctx, category)
	}
	return nil
}

func (m *MockService) DeleteCategory(ctx context.Context, id int) error {
	if m.DeleteCategoryFunc != nil {
		return m.DeleteCategoryFunc(ctx, id)
	}
	return nil
}

func (m *MockService) CreateProduct(ctx context.Context, product *models.Product) error {
	if m.CreateProductFunc != nil {
		// Asignar ID simulado
		product.ID = 1
		return m.CreateProductFunc(ctx, product)
	}
	product.ID = 1
	return nil
}

func (m *MockService) GetProduct(ctx context.Context, id int) (*models.Product, error) {
	if m.GetProductFunc != nil {
		return m.GetProductFunc(ctx, id)
	}
	return &models.Product{
		ID:       id,
		Name:     "Test Product",
		Price:    decimal.NewFromFloat(10.99),
		Stock:    100,
		IsActive: true,
	}, nil
}

func (m *MockService) GetAllProducts(ctx context.Context) ([]models.Product, error) {
	if m.GetAllProductsFunc != nil {
		return m.GetAllProductsFunc(ctx)
	}
	return []models.Product{
		{ID: 1, Name: "Product 1", Price: decimal.NewFromFloat(10.99), Stock: 100},
		{ID: 2, Name: "Product 2", Price: decimal.NewFromFloat(20.99), Stock: 50},
	}, nil
}

func (m *MockService) UpdateProduct(ctx context.Context, product *models.Product) error {
	if m.UpdateProductFunc != nil {
		return m.UpdateProductFunc(ctx, product)
	}
	return nil
}

func (m *MockService) DeleteProduct(ctx context.Context, id int) error {
	if m.DeleteProductFunc != nil {
		return m.DeleteProductFunc(ctx, id)
	}
	return nil
}

func (m *MockService) CreateCustomer(ctx context.Context, customer *models.Customer) error {
	if m.CreateCustomerFunc != nil {
		customer.ID = 1
		return m.CreateCustomerFunc(ctx, customer)
	}
	customer.ID = 1
	return nil
}

func (m *MockService) GetCustomer(ctx context.Context, id int) (*models.Customer, error) {
	if m.GetCustomerFunc != nil {
		return m.GetCustomerFunc(ctx, id)
	}
	return &models.Customer{
		ID:       id,
		FullName: "Test Customer",
		Email:    strPtr("test@example.com"),
	}, nil
}

func (m *MockService) GetAllCustomers(ctx context.Context) ([]models.Customer, error) {
	if m.GetAllCustomersFunc != nil {
		return m.GetAllCustomersFunc(ctx)
	}
	return []models.Customer{
		{ID: 1, FullName: "Customer 1"},
		{ID: 2, FullName: "Customer 2"},
	}, nil
}

func (m *MockService) UpdateCustomer(ctx context.Context, customer *models.Customer) error {
	if m.UpdateCustomerFunc != nil {
		return m.UpdateCustomerFunc(ctx, customer)
	}
	return nil
}

func (m *MockService) DeleteCustomer(ctx context.Context, id int) error {
	if m.DeleteCustomerFunc != nil {
		return m.DeleteCustomerFunc(ctx, id)
	}
	return nil
}

func (m *MockService) CreateSale(ctx context.Context, req *models.CreateSaleRequest) (*models.Sale, error) {
	if m.CreateSaleFunc != nil {
		return m.CreateSaleFunc(ctx, req)
	}
	return &models.Sale{
		ID:            1,
		CustomerID:    req.CustomerID,
		PaymentMethod: req.PaymentMethod,
		TotalAmount:   decimal.NewFromFloat(99.99),
	}, nil
}

func (m *MockService) GetSale(ctx context.Context, id int) (*models.Sale, error) {
	if m.GetSaleFunc != nil {
		return m.GetSaleFunc(ctx, id)
	}
	return &models.Sale{
		ID:          id,
		TotalAmount: decimal.NewFromFloat(99.99),
		Items: []models.SaleItem{
			{ID: 1, ProductID: 1, Quantity: 2, UnitPrice: decimal.NewFromFloat(49.995)},
		},
	}, nil
}

func (m *MockService) GetSalesReport(ctx context.Context) ([]models.SalesReportItem, error) {
	if m.GetSalesReportFunc != nil {
		return m.GetSalesReportFunc(ctx)
	}
	return []models.SalesReportItem{
		{Date: "2024-01-01", TotalSales: decimal.NewFromFloat(500.00)},
		{Date: "2024-01-02", TotalSales: decimal.NewFromFloat(750.00)},
	}, nil
}

func strPtr(s string) *string {
	return &s
}
