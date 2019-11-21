package middleware

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12/context"
)

// GetCORSpost returns CORS for any post request.
func GetCORSpost() context.Handler {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"HEAD", "GET", "POST", "DELETE"},
	})
}

// GetCORSdefault returns CORS for any thread or page request.
func GetCORSdefault() context.Handler {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"HEAD", "GET"},
	})
}
