package main

import (
	"meow-commerce-server/internal/product"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	// handle cors
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{
			"https://*", "http://*",
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Payment Intent Route
	mux.Get("/api/product", product.GetProducts)
	mux.Post("/api/product/{id}", product.GetProductById)
	mux.Get("/api/product/trending", product.GetTrendingProducts)

	return mux

}
