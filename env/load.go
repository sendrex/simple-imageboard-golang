package env

import (
	"fmt"

	"github.com/joho/godotenv"
)

func init() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		message := fmt.Errorf("[ENV FILE] Read FAILED @ %w", err)
		panic(message)
	} else {
		fmt.Println("[ENV FILE] Read OK")
	}
}
