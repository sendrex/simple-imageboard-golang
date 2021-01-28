package env

import (
	"fmt"

	"github.com/joho/godotenv"
)

func init() {
	// Read .env and panic if it couldn't be read
	if err := godotenv.Load(); err != nil {
		message := fmt.Errorf(".env file couldn't be read\n%w", err)
		panic(message)
	}
}
