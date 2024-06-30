package main

import (
	// "encoding/json"
	"meowcommerce-server/routes"
	"net/http"
)

func main() {

	// create multiplexer - recieves incoming requests and sorts them with URI patterns
	mux := http.NewServeMux()

	mux.Handle("/product/trending", &productHandler{})

	http.ListenAndServe(":4040", mux)
}
