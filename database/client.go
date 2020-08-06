package database

import (
	"fmt"

	"github.com/AquoDev/simple-imageboard-golang/env"
	"github.com/jinzhu/gorm"

	// For Postgres client
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Client connection to database.
var db *gorm.DB

func connect(settings string) (err error) {
	db, err = gorm.Open("postgres", settings)
	return
}

func init() {
	// Connect the client and panic if there are errors
	if err := connect(fmt.Sprintf(
		"host=%s port=%s dbname=%s sslmode=%s user=%s password=%s",
		env.GetString("DB_HOST"),
		env.GetString("DB_PORT"),
		env.GetString("DB_NAME"),
		env.GetString("DB_SSLMODE"),
		env.GetString("DB_USER"),
		env.GetString("DB_PASSWORD"),
	)); err != nil {
		message := fmt.Errorf("[DATABASE] Postgres connection failed @ %w", err)
		panic(message)
	}

	fmt.Println("[DATABASE] Postgres connection OK")
}
