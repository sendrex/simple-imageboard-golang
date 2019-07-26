package middleware

import "github.com/kataras/iris"

// UseContentTypeJSON sets the "Content-Type" header to "application/json; charset=utf-8".
func UseContentTypeJSON(ctx iris.Context) {
	ctx.Header("Content-Type", "application/json; charset=utf-8")
	ctx.Next()
}
