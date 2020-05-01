package database

import (
	"fmt"

	"github.com/AquoDev/simple-imageboard-golang/env"
	"github.com/jinzhu/gorm"

	// For Postgres client
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Client connection to database
var db *gorm.DB

func init() {
	// Parse connection settings
	settings := fmt.Sprintf(
		"host=%s port=%s dbname=%s sslmode=%s user=%s password=%s",
		env.GetString("DB_HOST"),
		env.GetString("DB_PORT"),
		env.GetString("DB_NAME"),
		env.GetString("DB_SSLMODE"),
		env.GetString("DB_USER"),
		env.GetString("DB_PASSWORD"),
	)

	// Connect the client and check if it's connected
	if conn, err := gorm.Open("postgres", settings); err != nil {
		message := fmt.Errorf("[DATABASE] Client connection FAILED @ %w", err)
		panic(message)
	} else {
		db = conn
		fmt.Println("[DATABASE] Client connection OK")
	}
}
