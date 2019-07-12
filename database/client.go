package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	// For Postgres client
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

// GetDatabaseClient returns an already connected DB client.
func GetDatabaseClient() *gorm.DB {
	return db
}

func init() {
	// Read values from .env file
	godotenv.Load()

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
		panic(err)
	} else {
		db = conn
	}
}
