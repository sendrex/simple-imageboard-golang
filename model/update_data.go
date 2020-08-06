package model

import "errors"

// UpdateData is a struct used for incoming PUT requests.
type UpdateData struct {
	ID       uint64 `json:"id"`
	Password string `json:"password"`
	Sticky   bool   `json:"sticky"`
	Closed   bool   `json:"closed"`
}

// Validate checks if all fields are valid.
func (data *UpdateData) Validate() error {
	if data.ID == 0 {
		return errors.New("id must be set")
	} else if data.Password == "" {
		return errors.New("password must not be empty")
	}

	return nil
}
