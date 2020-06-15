package env

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// GetString returns the string value from .env.
func GetString(key string) string {
	if env := os.Getenv(key); env == "" {
		message := fmt.Errorf("[ENV] %s env value is empty", key)
		panic(message)
	} else {
		return env
	}
}

// GetInt returns the int value from .env.
func GetInt(key string) int {
	if env, err := strconv.Atoi(GetString(key)); err != nil {
		message := fmt.Errorf("[ENV] Couldn't parse %s to int", key)
		panic(message)
	} else {
		return env
	}
}

// GetInt64 returns the int64 value from .env.
func GetInt64(key string) int64 {
	if env, err := strconv.ParseInt(GetString(key), 10, 64); err != nil {
		message := fmt.Errorf("[ENV] Couldn't parse %s to int64", key)
		panic(message)
	} else {
		return env
	}
}

// GetUint64 returns the uint64 value from .env.
func GetUint64(key string) uint64 {
	if env, err := strconv.ParseUint(GetString(key), 10, 64); err != nil {
		message := fmt.Errorf("[ENV] Couldn't parse %s to uint64", key)
		panic(message)
	} else {
		return env
	}
}

// GetTime returns the time.Duration value from .env.
func GetTime(key string, magnitude string) time.Duration {
	timeString := fmt.Sprintf("%s%s", GetString(key), magnitude)
	if env, err := time.ParseDuration(timeString); err != nil {
		message := fmt.Errorf("[ENV] Couldn't parse %s to time.Duration (%s)", key, magnitude)
		panic(message)
	} else {
		return env
	}
}
