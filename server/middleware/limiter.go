package middleware

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/AquoDev/simple-imageboard-golang/server/handler"
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

func configureLimiter(requestPerSecond float64, windowDuration time.Duration) (lmt *limiter.Limiter) {
	var message interface{}
	message = handler.GetError(429)
	message, _ = json.Marshal(message)
	jsonMessage := string(message.([]byte))

	lmt = tollbooth.NewLimiter(requestPerSecond, nil)
	lmt.SetTokenBucketExpirationTTL(windowDuration * time.Second)
	lmt.SetMessage(jsonMessage)
	lmt.SetMessageContentType("application/json; charset=UTF-8")

	return
}

func init() {
	regular = configureLimiter(3, 15)
	strict = configureLimiter(0.1, 15)
	fmt.Println("Limiters are now configured")
}
