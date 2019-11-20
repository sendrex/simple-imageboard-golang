package middleware

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/AquoDev/simple-imageboard-golang/handler"
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
)

var (
	regular, strict *limiter.Limiter
)

// RegularLimiter returns an already configured limiter.
func RegularLimiter() *limiter.Limiter {
	return regular
}

// StrictLimiter returns an already configured limiter.
func StrictLimiter() *limiter.Limiter {
	return strict
}

func configureLimiter(requestsPerSecond float64) *limiter.Limiter {
	message, err := json.Marshal(handler.GetError(429))
	if err != nil {
		panic(err)
	}

	return tollbooth.NewLimiter(requestsPerSecond, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour}).SetMessage(string(message)).SetMessageContentType("application/json; charset=UTF-8").SetStatusCode(200)
}

func init() {
	regular = configureLimiter(3)
	strict = configureLimiter(0.1)
	fmt.Println("Limiters: OK")
}
