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
		fmt.Println("[DATABASE]: Postgres client OK")
	}
}
