package product

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Products
func GetProducts(w http.ResponseWriter, r *http.Request) {
	// get data from database
	products := GetAllProducts()

	// convert products to only productSummaries for list showcase
	var productSummaryList []ProductSummary

	for _, product := range products {
		productSummaryList = append(productSummaryList, ProductSummary{
			ID:       product.ID,
			Title:    product.Title,
			Subtitle: product.Subtitle,
			ImageUrl: product.ImageUrl,
		})
	}

	out, err := json.Marshal(productSummaryList)

	if err != nil {
		fmt.Println("Error when attempting to marshal json:", productSummaryList)
	}

	// set headers
	w.Header().Set("Content-Type", "application/json")

	// return response
	w.Write(out)
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
	products := GetAllProducts()

	// convert products to only productSummaries for list showcase
	var productSummaryList []ProductSummary

	for _, product := range products {
		productSummaryList = append(productSummaryList, ProductSummary{
			ID:       product.ID,
			Title:    product.Title,
			Subtitle: product.Subtitle,
			ImageUrl: product.ImageUrl,
		})
	}

	out, err := json.Marshal(productSummaryList)

	if err != nil {
		fmt.Println("Error when attempting to marshal json:", productSummaryList)
	}

	// set headers
	w.Header().Set("Content-Type", "application/json")

	// return response
	w.Write(out)
}
