package server

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
	"web-server/internal/database"
	"web-server/internal/handlers"
	"web-server/internal/middleware"
)

type Server struct {
	port int
	db   database.Service
}

func NewServer(port int, db database.Service) *http.Server {
	s := &Server{
		port: port,
		db:   db,
	}

	// Chain middlewares
	handler := s.registerRoutes()
	handler = middleware.Logger(handler)
	handler = middleware.Recovery(handler)
	handler = middleware.RequestID(handler)
	handler = middleware.Timeout(30 * time.Second)(handler)

	// CORS from env
	corsOrigins := getCORSOrigins()
	handler = middleware.CORS(corsOrigins)(handler)

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", s.port),
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

func (s *Server) registerRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.HealthCheck(s.db))

	// Categories
	categoryHandler := handlers.NewCategoryHandler(s.db)
	mux.HandleFunc("POST /categories", categoryHandler.CreateCategory)
	mux.HandleFunc("GET /categories/{id}", categoryHandler.GetCategory)
	mux.HandleFunc("GET /categories", categoryHandler.GetAllCategories)
	mux.HandleFunc("PUT /categories/{id}", categoryHandler.UpdateCategory)
	mux.HandleFunc("DELETE /categories/{id}", categoryHandler.DeleteCategory)

	// Products
	productHandler := handlers.NewProductHandler(s.db)
	mux.HandleFunc("POST /products", productHandler.CreateProduct)
	mux.HandleFunc("GET /products/{id}", productHandler.GetProduct)
	mux.HandleFunc("GET /products", productHandler.GetAllProducts)
	mux.HandleFunc("PUT /products/{id}", productHandler.UpdateProduct)
	mux.HandleFunc("DELETE /products/{id}", productHandler.DeleteProduct)

	// Customers
	customerHandler := handlers.NewCustomerHandler(s.db)
	mux.HandleFunc("POST /customers", customerHandler.CreateCustomer)
	mux.HandleFunc("GET /customers/{id}", customerHandler.GetCustomer)
	mux.HandleFunc("GET /customers", customerHandler.GetAllCustomers)
	mux.HandleFunc("PUT /customers/{id}", customerHandler.UpdateCustomer)
	mux.HandleFunc("DELETE /customers/{id}", customerHandler.DeleteCustomer)

	// Sales
	saleHandler := handlers.NewSaleHandler(s.db)
	mux.HandleFunc("POST /sales", saleHandler.CreateSale)
	mux.HandleFunc("GET /sales/report", saleHandler.GetSalesReport)
	mux.HandleFunc("GET /sales/{id}", saleHandler.GetSale)

	return mux
}

func getCORSOrigins() []string {
	origins := os.Getenv("CORS_ORIGINS")
	if origins == "" {
		// Default: allow all in development
		return []string{"*"}
	}
	return strings.Split(origins, ",")
}
