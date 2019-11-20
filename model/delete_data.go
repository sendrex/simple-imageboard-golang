package model

// DeleteData is a struct in which data to delete a post is parsed from JSON.
type DeleteData struct {
	ID         uint64 `json:"id"`
	DeleteCode string `json:"delete_code"`
}
