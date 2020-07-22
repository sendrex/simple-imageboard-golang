package middleware

import (
	"github.com/AquoDev/simple-imageboard-golang/framework"
)

const expectedContentType = "application/json; charset=UTF-8"

// CheckHeader handles a response with an error if the header isn't valid.
func CheckHeader() framework.MiddlewareFunc {
	return framework.CheckHeader(expectedContentType)
}
