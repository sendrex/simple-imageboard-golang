package handler

import (
	"net/http"

	"github.com/AquoDev/simple-imageboard-golang/framework"
)

var healthMessage = map[string]string{
	"message": http.StatusText(http.StatusOK),
}

// GetHealthcheck handles a JSON response with a 200 OK if everything's alright.
func GetHealthcheck(ctx framework.Context) error {
	return framework.SendOK(ctx, healthMessage)
}
