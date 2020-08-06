package model

import "errors"

// DeleteData is a struct used for incoming DELETE requests or successful POST requests.
type DeleteData struct {
	ID       uint64 `json:"id"`
	Password string `json:"password"`
}

// Validate checks if all fields are valid.
func (data *DeleteData) Validate() error {
	if data.ID == 0 {
		return errors.New("id must be set")
	} else if data.Password == "" {
		return errors.New("password must not be empty")
	}

	return nil
}
