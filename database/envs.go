package database

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	threadsPerPage, postsPerThread, maxRootThreads uint64
)

func init() {
	// Read values from .env file
	godotenv.Load()

	// Parse board settings
	if parseInt, err := strconv.ParseUint(os.Getenv("THREADS_PER_PAGE"), 10, 0); err != nil {
		panic(err)
	} else {
		threadsPerPage = parseInt
	}

	if parseInt, err := strconv.ParseUint(os.Getenv("POSTS_PER_THREAD"), 10, 0); err != nil {
		panic(err)
	} else {
		postsPerThread = parseInt
	}

	maxRootThreads = threadsPerPage * 10

	fmt.Println("[DATABASE]: Board settings OK")
}
