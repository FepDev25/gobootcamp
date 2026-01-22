package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
	"web-server/internal/database"
	"web-server/internal/models"

	"github.com/joho/godotenv"
	"github.com/shopspring/decimal"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	dbConnStr := os.Getenv("DB_CONN_STR")
	if dbConnStr == "" {
		log.Fatal("DB_CONN_STR environment variable is not set")
	}

	// Initialize DB
	db, err := database.New(dbConnStr)
	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}
	defer db.Close()

	ctx := context.Background()

	log.Println("Starting seeding process...")

	// 1. Categories
	categories := []models.Category{
		{Name: "Bebidas", Description: "Todo tipo de bebidas"},
		{Name: "Lácteos", Description: "Leche, quesos y derivados"},
		{Name: "Panadería", Description: "Pan fresco y facturas"},
		{Name: "Limpieza", Description: "Productos para el hogar"},
		{Name: "Carnicería", Description: "Carnes rojas y blancas"},
		{Name: "Verdulería", Description: "Frutas y verduras frescas"},
		{Name: "Snacks", Description: "Papas fritas, chizitos, etc"},
		{Name: "Almacén", Description: "Productos generales"},
	}

	createdCategories := make([]int, 0)
	for _, c := range categories {
		// Try to create (ignore error if exists for unique name constraint in simple seed)
		err := db.CreateCategory(ctx, &c)
		if err == nil {
			fmt.Printf("Created Category: %s (ID: %d)\n", c.Name, c.ID)
			createdCategories = append(createdCategories, c.ID)
		} else {
			fmt.Printf("Skipping category %s (probably exists)\n", c.Name)
			// Try to find it if you want to be robust, but for seeding fresh DB it's fine.
			// In a real robust seeder we would fetch the ID.
			// For now, let's assume if it fails we might not have the ID easily without a GetByName method.
			// Let's just fetch all categories to populate our list if the insert failed.
		}
	}

	// Refresh category IDs in case some already existed
	allCats, err := db.GetAllCategories(ctx)
	if err != nil {
		log.Fatalf("Failed to fetch categories: %v", err)
	}
	createdCategories = nil
	for _, c := range allCats {
		createdCategories = append(createdCategories, c.ID)
	}

	if len(createdCategories) == 0 {
		log.Fatal("No categories available to create products")
	}

	// 2. Products
	productNames := []string{"Coca Cola", "Pepsi", "Leche La Serenísima", "Pan Lactal", "Detergente Ala", "Jabón Dove", "Asado", "Vacío", "Tomate", "Lechuga", "Papas Lays", "Galletitas Oreos", "Arroz Gallo", "Fideos Matarazzo", "Aceite Natura"}
	createdProducts := make([]int, 0)

	for i := 0; i < 50; i++ {
		name := fmt.Sprintf("%s %d", productNames[rand.Intn(len(productNames))], i)
		catID := createdCategories[rand.Intn(len(createdCategories))]
		barcode := fmt.Sprintf("BARCODE-%d-%d", time.Now().UnixNano(), i)
		price := decimal.NewFromFloat(10.0 + rand.Float64()*90.0).Round(2) // 10.00 to 100.00
		stock := 100 + rand.Intn(900)

		p := models.Product{
			CategoryID: &catID,
			Barcode:    &barcode,
			Name:       name,
			Price:      price,
			Stock:      stock,
			IsActive:   true,
		}

		err := db.CreateProduct(ctx, &p)
		if err != nil {
			log.Printf("Failed to create product %s: %v", name, err)
		} else {
			// fmt.Printf("Created Product: %s (ID: %d)\n", p.Name, p.ID)
			createdProducts = append(createdProducts, p.ID)
		}
	}
	fmt.Printf("Created %d products\n", len(createdProducts))

	// 3. Customers
	customerNames := []string{"Juan Perez", "Maria Garcia", "Carlos Lopez", "Ana Martinez", "Jose Gonzalez", "Laura Rodriguez", "Miguel Fernandez", "Sofia Lopez"}
	createdCustomers := make([]int, 0)

	for i := 0; i < 20; i++ {
		name := fmt.Sprintf("%s %d", customerNames[rand.Intn(len(customerNames))], i)
		email := fmt.Sprintf("user%d_%d@example.com", i, time.Now().UnixNano())
		phone := fmt.Sprintf("+54911%d", rand.Intn(99999999))
		addr := fmt.Sprintf("Calle Falsa %d", rand.Intn(1000))

		c := models.Customer{
			FullName: name,
			Email:    &email,
			Phone:    &phone,
			Address:  &addr,
		}

		err := db.CreateCustomer(ctx, &c)
		if err != nil {
			log.Printf("Failed to create customer %s: %v", name, err)
		} else {
			// fmt.Printf("Created Customer: %s (ID: %d)\n", c.FullName, c.ID)
			createdCustomers = append(createdCustomers, c.ID)
		}
	}
	fmt.Printf("Created %d customers\n", len(createdCustomers))

	if len(createdProducts) == 0 || len(createdCustomers) == 0 {
		log.Println("Not enough data to create sales")
		return
	}

	// 4. Sales
	log.Println("Creating random sales...")
	for i := 0; i < 50; i++ {
		custID := createdCustomers[rand.Intn(len(createdCustomers))]

		// 1 to 5 items per sale
		numItems := 1 + rand.Intn(5)
		items := make([]models.CreateSaleItemRequest, 0, numItems)

		// Avoid duplicate products in same sale for simplicity, or allow it
		for j := 0; j < numItems; j++ {
			prodID := createdProducts[rand.Intn(len(createdProducts))]
			qty := 1 + rand.Intn(5)

			items = append(items, models.CreateSaleItemRequest{
				ProductID: prodID,
				Quantity:  qty,
			})
		}

		req := models.CreateSaleRequest{
			CustomerID:    &custID,
			PaymentMethod: "CASH", // or random between CASH, CARD, QR
			Items:         items,
		}

		_, err := db.CreateSale(ctx, &req)
		if err != nil {
			log.Printf("Failed to create sale: %v", err)
		} else {
			// fmt.Printf("Created Sale ID: %d\n", sale.ID)
		}
	}
	fmt.Println("Seeding completed successfully!")
}
