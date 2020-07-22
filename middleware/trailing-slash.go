package middleware

import (
	"github.com/AquoDev/simple-imageboard-golang/framework"
)

// RemoveTrailingSlash removes the trailing slash from URLs.
func RemoveTrailingSlash() framework.MiddlewareFunc {
	return framework.RemoveTrailingSlash()
}
