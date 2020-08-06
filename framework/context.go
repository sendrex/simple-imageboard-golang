package framework

import (
	"net/http"
	"strconv"

	"github.com/AquoDev/simple-imageboard-golang/model"
	"github.com/labstack/echo/v4"
)

// SendOK sends a JSON 200 OK response with any given data as the body.
func SendOK(ctx Context, data interface{}) error {
	return ctx.JSON(http.StatusOK, data)
}

// SendCreated sends a JSON 201 Created response with a delete code struct as the body.
func SendCreated(ctx Context, data *model.DeleteData) error {
	return ctx.JSON(http.StatusCreated, data)
}

// SendError sends any error and a message as the body.
func SendError(status int) *HTTPError {
	return echo.NewHTTPError(status)
}

// GetID gets the ID param from the URL and is returned with an error to be checked in the handler.
func GetID(ctx Context) (uint64, error) {
	return strconv.ParseUint(ctx.Param("id"), 10, 64)
}

// BindPost returns a post struct extracted from the body and an error to be checked in the handler.
func BindPost(ctx Context) (*model.Post, error) {
	post := new(model.Post)

	if err := ctx.Bind(&post); err != nil {
		// If data couldn't be binded, return that error
		return nil, err
	}

	return post, post.Validate()
}

// BindDeleteData returns a delete data struct extracted from the body and an error to be checked in the handler.
func BindDeleteData(ctx Context) (*model.DeleteData, error) {
	deleteData := new(model.DeleteData)

	if err := ctx.Bind(&deleteData); err != nil {
		// If data couldn't be binded, return that error
		return nil, err
	}

	return deleteData, deleteData.Validate()
}

// BindUpdateData returns a update data struct extracted from the body and an error to be checked in the handler.
func BindUpdateData(ctx Context) (*model.UpdateData, error) {
	updateData := new(model.UpdateData)

	if err := ctx.Bind(&updateData); err != nil {
		// If data couldn't be binded, return that error
		return nil, err
	}

	return updateData, updateData.Validate()
}
