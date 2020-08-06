package env

import (
	"fmt"

	"github.com/joho/godotenv"
)

func init() {
	// Read .env and panic if it couldn't be read
	if err := godotenv.Load(); err != nil {
		message := fmt.Errorf("[ENV] file .env couldn't be read @ %w", err)
		panic(message)
	}

	fmt.Println("[ENV] file .env read OK")
}
