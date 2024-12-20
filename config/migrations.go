package config

import (
	"github.com/jmoiron/sqlx"
	"log"
)

/**
* Installs and uses uuid extension for *postgres*.
* NOTE: !!!Currently using goose and not this file for migrations!!!
**/
func SetupUUIDExtension(db *sqlx.DB) {
	query := `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create uuid-ossp extension: %v", err)
	}
	log.Println("UUID extension setup completed.")
}

/**
* Sets up all migration steps for all tables.
**/
func RunMigrations(db *sqlx.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
        name TEXT NOT NULL,
        email TEXT UNIQUE NOT NULL,
				password TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

    `

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	log.Println("Migrations ran successfully.")
}
