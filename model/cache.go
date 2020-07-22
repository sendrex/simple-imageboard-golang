package model

// Cache is a struct used for caching posts, threads, index or errors.
type Cache struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}
