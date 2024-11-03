package main

import (
	"fmt"
	"log"
	"meow-commerce-server/internal/config"
	"meow-commerce-server/internal/driver"
	"net/http"
	"time"

	"github.com/joho/godotenv"
)

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
	dsn := "root:12456@tcp(localhost:3308)/virtual_terminal_db?parseTime=true&tls=false"

	dbCfg := config.LoadDBConfig()

	cfg := config.LoadConfig(4040, dbCfg)
	// dbCfg :=

	// initialize app
	app := &application{
		config: cfg,
	}

	// initialize db connection
	_, err := driver.OpenDB(app.config.db.dsn)

	if err != nil {
		// print DB error
		log.Fatalf("Could not connect to DB:", err)
	}

	if err := app.serve();

	// check if error was set
	err != nil {
		// throw error
		log.Fatalf("Could not start server %v", err)
	}
}
