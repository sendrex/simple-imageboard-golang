package env

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// GetString returns the string value from .env.
func GetString(key string) string {
	env := os.Getenv(key)
	if env == "" {
		message := fmt.Errorf("%s environment variable is empty (and it must not be)", key)
		panic(message)
	}
	return env
}

// GetInt returns the int value from .env.
func GetInt(key string) int {
	env, err := strconv.Atoi(GetString(key))
	if err != nil {
		message := fmt.Errorf("GetInt() failed, couldn't parse %s to int\n%w", key, err)
		panic(message)
	}
	return env
}

// GetInt64 returns the int64 value from .env.
func GetInt64(key string) int64 {
	env, err := strconv.ParseInt(GetString(key), 10, 64)
	if err != nil {
		message := fmt.Errorf("GetInt64() failed, couldn't parse %s to int64\n%w", key, err)
		panic(message)
	}
	return env
}

// GetUint64 returns the uint64 value from .env.
func GetUint64(key string) uint64 {
	env, err := strconv.ParseUint(GetString(key), 10, 64)
	if err != nil {
		message := fmt.Errorf("GetUint64() failed, couldn't parse %s to uint64\n%w", key, err)
		panic(message)
	}
	return env
}

// GetTime returns the time.Duration value from .env.
func GetTime(key string, magnitude string) time.Duration {
	timeString := fmt.Sprintf("%s%s", GetString(key), magnitude)
	env, err := time.ParseDuration(timeString)
	if err != nil {
		message := fmt.Errorf("GetTime() failed, couldn't parse %s to time.Duration (%s)\n%w", key, magnitude, err)
		panic(message)
	}
	return env
}
