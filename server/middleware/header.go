package middleware

import (
	"strings"

	"github.com/AquoDev/simple-imageboard-golang/server/handler"
	"github.com/kataras/iris"
)

// UseContentTypeJSON sets the "Content-Type" header to "application/json; charset=utf-8".
func UseContentTypeJSON(ctx iris.Context) {
	ctx.Header("Content-Type", "application/json; charset=utf-8")
	ctx.Next()
}

// CheckHeaders handles a response with an error if any check fails.
func CheckHeaders(ctx iris.Context) {
	contentType := ctx.GetHeader("Content-Type")

	// Check "Content-Type"
	if strings.ToLower(contentType) != "application/json; charset=utf-8" {
		invalidHeader := handler.GetError(400)
		ctx.WriteString(invalidHeader)
		return
	}

	ctx.Next()
}
