package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"web-server/internal/database"
	"web-server/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	// Load env vars
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	if port == 0 {
		port = 8080
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

	srv := server.NewServer(port, db)

	fmt.Printf("Server starting on port %d...\n", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("cannot start server: %v", err)
	}
}
