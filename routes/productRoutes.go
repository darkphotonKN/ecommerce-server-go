package routes

import (
	"meowcommerce-server/models"
	"net/http"
)

type productHandler struct{}

func (h *productHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("These are the products."))
}
