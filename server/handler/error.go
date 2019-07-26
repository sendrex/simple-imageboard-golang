package handler

import (
	"encoding/json"

	"github.com/kataras/iris"
)

type errorJSON struct {
	Status  uint16 `json:"status"`
	Message string `json:"message"`
}

var (
	responses = make(map[uint16]string)
	errors    = map[uint16]string{
		400: "Bad Request",
		404: "Not Found",
		500: "Internal Server Error",
	}
)

// GetError returns a JSON string with a <status> <message> response.
func GetError(status uint16) (response string) {
	if value, ok := responses[status]; ok {
		response = value
	} else {
		response = GetError(500)
	}
	return
}

// PathNotFound handles a failed request with a 404 Not Found response.
func PathNotFound(ctx iris.Context) {
	response := GetError(404)
	ctx.WriteString(response)
}

func makeResponseError(status uint16, message string) string {
	jsonObject, _ := json.Marshal(&errorJSON{
		Status:  status,
		Message: message,
	})
	return string(jsonObject)
}

func init() {
	for status, message := range errors {
		responses[status] = makeResponseError(status, message)
	}
}
