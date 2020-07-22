package middleware

import (
	"github.com/AquoDev/simple-imageboard-golang/framework"
)

// Secure adds to the response some headers to increase protection.
func Secure() framework.MiddlewareFunc {
	return framework.Secure()
}
