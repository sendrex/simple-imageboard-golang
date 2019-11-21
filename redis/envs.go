package redis

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var (
	maxTimePage, maxTimeThread, maxTimePost time.Duration
)

func init() {
	// Read values from .env file
	godotenv.Load()

	// Parse every expiry time
	if parsedTime, err := time.ParseDuration(os.Getenv("REDIS_EXPIRE_TIME_PAGE")); err != nil {
		panic(err)
	} else {
		maxTimePage = parsedTime
	}

	if parsedTime, err := time.ParseDuration(os.Getenv("REDIS_EXPIRE_TIME_THREAD")); err != nil {
		panic(err)
	} else {
		maxTimeThread = parsedTime
	}

	if parsedTime, err := time.ParseDuration(os.Getenv("REDIS_EXPIRE_TIME_POST")); err != nil {
		panic(err)
	} else {
		maxTimePost = parsedTime
	}

	fmt.Println("All Redis expiry times parsed successfully")
}
