package main

import (
	"fmt"
	"log"
	"os"

	"github.com/darkphotonKN/ecommerce-server-go/config"
	seeder "github.com/darkphotonKN/ecommerce-server-go/internal/utils"
	"github.com/joho/godotenv"
)

/**
* Main entry point to entire application.
* NOTE: Keep code here as clean and little as possible.
**/
func main() {
	// env setup
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// database setup
	db := config.InitDB()
	defer db.Close()

	// router setup
	router := config.SetupRouter()

	defaultDevPort := ":8080"

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultDevPort
	}

	// TODO: Remove for production.
	// Seeding default products for testing.
	seeder.SeedProducts()

	// starts server and listen on port
	router.Run(fmt.Sprintf(":%s", port)) // port = ":" + PORT
}
