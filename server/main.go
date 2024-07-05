package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

type application struct {
	config config
}

func (app *application) serve() error {

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	fmt.Println("Server running on", app.config.port)
	return srv.ListenAndServe()
}

func main() {
	config := config{
		port: 4040,
	}

	// initialize app
	app := &application{
		config: config,
	}

	if err := app.serve();

	// check if error was set
	err != nil {
		// throw error
		log.Fatalf("Could not start server %v", err)
	}
}
