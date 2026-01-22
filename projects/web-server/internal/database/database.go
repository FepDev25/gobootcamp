package database

import (
	"context"
	"fmt"
	"time"
	"web-server/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Service interface {
	Health() map[string]string
	Close()
	GetPool() *pgxpool.Pool

	// Categories
	CreateCategory(ctx context.Context, category *models.Category) error
	GetCategory(ctx context.Context, id int) (*models.Category, error)
	GetAllCategories(ctx context.Context) ([]models.Category, error)
	UpdateCategory(ctx context.Context, category *models.Category) error
	DeleteCategory(ctx context.Context, id int) error

	// Products
	CreateProduct(ctx context.Context, product *models.Product) error
	GetProduct(ctx context.Context, id int) (*models.Product, error)
	GetAllProducts(ctx context.Context) ([]models.Product, error)
	UpdateProduct(ctx context.Context, product *models.Product) error
	DeleteProduct(ctx context.Context, id int) error

	// Customers
	CreateCustomer(ctx context.Context, customer *models.Customer) error
	GetCustomer(ctx context.Context, id int) (*models.Customer, error)
	GetAllCustomers(ctx context.Context) ([]models.Customer, error)
	UpdateCustomer(ctx context.Context, customer *models.Customer) error
	DeleteCustomer(ctx context.Context, id int) error

	// Sales
	CreateSale(ctx context.Context, req *models.CreateSaleRequest) (*models.Sale, error)
	GetSale(ctx context.Context, id int) (*models.Sale, error)
}

type service struct {
	db *pgxpool.Pool
}

var (
	dbInstance *service
)

func New(connStr string) (Service, error) {
	// Reusing instance if it exists (Singleton pattern logic if strictly needed,
	// but standard dependency injection in main is preferred. Keeping it simple).
	if dbInstance != nil {
		return dbInstance, nil
	}

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("unable to parse connection string: %w", err)
	}

	// Basic connection pool config
	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnLifetime = time.Hour

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	// Ping to verify connection
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	dbInstance = &service{
		db: pool,
	}

	return dbInstance, nil
}

func (s *service) GetPool() *pgxpool.Pool {
	return s.db
}

func (s *service) Close() {
	s.db.Close()
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// Ping the DB
	err := s.db.Ping(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		return stats
	}

	// Database status
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Pool stats
	dbStats := s.db.Stat()
	stats["open_connections"] = fmt.Sprintf("%d", dbStats.TotalConns())
	stats["in_use"] = fmt.Sprintf("%d", dbStats.AcquiredConns())
	stats["idle"] = fmt.Sprintf("%d", dbStats.IdleConns())

	return stats
}
