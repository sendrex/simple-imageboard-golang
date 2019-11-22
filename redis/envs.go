package redis

import (
	"fmt"
	"os"
	"time"
)

var (
	maxTimePage, maxTimeThread, maxTimePost time.Duration
)

func init() {
	stringTimePage := fmt.Sprintf("%ss", os.Getenv("REDIS_EXPIRE_TIME_PAGE"))
	stringTimeThread := fmt.Sprintf("%ss", os.Getenv("REDIS_EXPIRE_TIME_THREAD"))
	stringTimePost := fmt.Sprintf("%ss", os.Getenv("REDIS_EXPIRE_TIME_POST"))

	// Parse every expiry time
	if parsedTime, err := time.ParseDuration(stringTimePage); err != nil {
		panic(err)
	} else {
		maxTimePage = parsedTime
	}

	if parsedTime, err := time.ParseDuration(stringTimeThread); err != nil {
		panic(err)
	} else {
		maxTimeThread = parsedTime
	}

	if parsedTime, err := time.ParseDuration(stringTimePost); err != nil {
		panic(err)
	} else {
		maxTimePost = parsedTime
	}

	fmt.Println("[REDIS]: Expiration times OK")
}
