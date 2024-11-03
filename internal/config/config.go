package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type db struct {
	cfg dbConfig
}

type config struct {
	port int
	env  string
	db   db
}

type dbConfig struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
	Dsn        string
}

func LoadDBConfig(dsn string) *dbConfig {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConfig := dbConfig{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		Dsn:        dsn,
	}

	return &dbConfig

}

func LoadConfig(port int, db dbConfig) *config {

	return &config{
		port: port,
		db:   db,
	}
}
