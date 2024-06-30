package main

import (
	// "encoding/json"
	"log"
	"meow-commerce-server/internal/product"
	"net/http"
)

func main() {
	// Products Route
	http.HandleFunc("/api/products", product.GetProducts)
	http.HandleFunc("/api/products/trending", product.GetTrendingProducts)

	log.Println("Starting server on 4040.")

	if err := http.ListenAndServe(":4040", nil);
	// check if error was set
	err != nil {
		// throw error
		log.Fatalf("Could not start server %v", err)

	}
}
