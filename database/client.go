package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	// For Postgres client
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Client connection to database
var db *gorm.DB

func init() {
	// Load .env for the first time
	if err := godotenv.Load(); err != nil {
		message := fmt.Errorf("[ENV FILE] Read FAILED @ %w", err)
		panic(message)
	} else {
		fmt.Println("[ENV FILE] Read OK")
	}

	// Parse connection settings
	settings := fmt.Sprintf(
		"host=%s port=%s dbname=%s sslmode=%s user=%s password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
	)

	// Connect the client
	conn, err := gorm.Open("postgres", settings)
	if err != nil {
		message := fmt.Errorf("[DATABASE] Client connection FAILED @ %w", err)
		panic(message)
	} else {
		db = conn
		fmt.Println("[DATABASE] Client connection OK")
	}
}
