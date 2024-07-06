package product

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Products
func GetProducts(w http.ResponseWriter, r *http.Request) {
	// get data from database
	products := GetAllProducts()

	// set headers
	w.Header().Set("Content-Type", "application/json")

	// encode products into json and write, checking if there is an error
	json.NewEncoder(w).Encode(products)
}

// Get Single Product with ID
func GetProductById(w http.ResponseWriter, r *http.Request) {
	// get the product Id
	productId := chi.URLParam(r, "id")

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
		id, _ := strconv.ParseInt(productId, 10, 64)

		if int64(product.ID) == id {
			productFound = product

			// stop redundant loops
			break
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
	// get data from database
	w.Header().Set("Content-Type", "application/json")

	products := GetAllTrendingProducts()

	json.NewEncoder(w).Encode(products)
}
