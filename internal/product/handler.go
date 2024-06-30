package product

import (
	"encoding/json"
	"net/http"
)

// Products
func GetProducts(w http.ResponseWriter, r *http.Request) {
	// handle incorrect method handling
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
	}

	// get data from database
	products := GetAllProducts()

	// set headers
	w.Header().Set("Content-Type", "application/json")

	// encode products into json and write, checking if there is an error
	json.NewEncoder(w).Encode(products)
}

// Trending Products
func GetTrendingProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed for Get Trending Products", http.StatusNotFound)
	}

	// get data from database
	w.Header().Set("Content-Type", "application/json")

	products := GetAllTrendingProducts()

	json.NewEncoder(w).Encode(products)
}
