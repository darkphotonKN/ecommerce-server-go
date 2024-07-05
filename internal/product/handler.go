package product

import (
	"encoding/json"
	"log"
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

// Get Single Product with ID
func GetProductById(w http.ResponseWriter, r *http.Request) {
	var productId ProductId
	// get body from request
	body := r.Body
	err := json.NewDecoder(body).Decode(&productId)

	if err != nil {
		log.Println("Error when decoding json.")
	}

	products := GetAllProducts()

	// loop and return product
	var productFound Product
	for _, product := range products {
		if product.ID == productId.ID {
			productFound = product
		}
	}

	out, err := json.Marshal(productFound)

	if err != nil {
		log.Println("Error when marshalled json.")
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(out)
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
