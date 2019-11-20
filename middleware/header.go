package middleware

import (
	"github.com/AquoDev/simple-imageboard-golang/handler"
	"github.com/kataras/iris/v12"
)

// CheckHeaders handles a response with an error if any check fails.
func CheckHeaders(ctx iris.Context) {
	// Check "Content-Type"
	if ctx.GetHeader("Content-Type") != "application/json; charset=UTF-8" {
		invalidHeader := handler.GetError(400)
		ctx.JSON(invalidHeader)
		return
	}

	ctx.Next()
}
