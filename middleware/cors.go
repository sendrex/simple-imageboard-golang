package middleware

import (
	"net/http"

	"github.com/AquoDev/simple-imageboard-golang/framework"
)

// ExtendedCORS allows GET, POST and DELETE requests.
func ExtendedCORS() framework.MiddlewareFunc {
	return framework.CORS(http.MethodOptions, http.MethodGet, http.MethodHead, http.MethodPost, http.MethodDelete)
}

// DefaultCORS allows only GET requests.
func DefaultCORS() framework.MiddlewareFunc {
	return framework.CORS(http.MethodOptions, http.MethodGet, http.MethodHead)
}
