package handler

import (
	"github.com/kataras/iris/v12"
)

var errors = map[uint16]string{
	400: "Bad Request",
	404: "Not Found",
	429: "Too Many Requests",
	500: "Internal Server Error",
}

// GetError returns a map structure with a <status> <message> response.
func GetError(status uint16) (response map[string]interface{}) {
	if message, ok := errors[status]; ok {
		response = map[string]interface{}{
			"message": message,
			"status":  status,
		}
	} else {
		response = GetError(500)
	}
	return
}

// PathNotFound handles a failed request with a 404 Not Found response.
func PathNotFound(ctx iris.Context) {
	response := GetError(404)
	ctx.JSON(response)
}
