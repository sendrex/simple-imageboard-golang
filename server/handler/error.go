package handler

import (
	"encoding/json"
	"github.com/kataras/iris"
)

type errorJSON struct {
	Status  uint16 `json:"status"`
	Message string `json:"message"`
}

var notFoundResponse string

// GetNotFoundResponse returns a JSON string with a 404 Not Found response.
func GetNotFoundResponse() string {
	return notFoundResponse
}

// PathNotFound handles a failed request.
func PathNotFound(ctx iris.Context) {
	ctx.WriteString(notFoundResponse)
}

func makeResponseError(status uint16, message string) string {
	jsonObject, _ := json.Marshal(&errorJSON{
		Status:  status,
		Message: message,
	})

	return string(jsonObject)
}

func init() {
	notFoundResponse = makeResponseError(404, "Not Found")
}
