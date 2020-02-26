package database

import (
	"os"
	"strconv"
)

var (
	threadsPerPage, postsPerThread, maxRootThreads uint64
)

func init() {
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
}
