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
		message := fmt.Errorf("[ENV] %s value is empty", key)
		panic(message)
	}
	return env
}

// GetInt returns the int value from .env.
func GetInt(key string) int {
	env, err := strconv.Atoi(GetString(key))
	if err != nil {
		message := fmt.Errorf("[ENV] GetInt() failed, couldn't parse %s to int @ %w", key, err)
		panic(message)
	}
	return env
}

// GetInt64 returns the int64 value from .env.
func GetInt64(key string) int64 {
	env, err := strconv.ParseInt(GetString(key), 10, 64)
	if err != nil {
		message := fmt.Errorf("[ENV] GetInt64() failed, couldn't parse %s to int64 @ %w", key, err)
		panic(message)
	}
	return env
}

// GetUint64 returns the uint64 value from .env.
func GetUint64(key string) uint64 {
	env, err := strconv.ParseUint(GetString(key), 10, 64)
	if err != nil {
		message := fmt.Errorf("[ENV] GetUint64() failed, couldn't parse %s to uint64 @ %w", key, err)
		panic(message)
	}
	return env
}

// GetTime returns the time.Duration value from .env.
func GetTime(key string, magnitude string) time.Duration {
	timeString := fmt.Sprintf("%s%s", GetString(key), magnitude)
	env, err := time.ParseDuration(timeString)
	if err != nil {
		message := fmt.Errorf("[ENV] GetTime() failed, couldn't parse %s to time.Duration (%s) @ %w", key, magnitude, err)
		panic(message)
	}
	return env
}
