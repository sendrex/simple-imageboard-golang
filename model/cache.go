package model

// Cache is a struct in which data to cache with a status is parsed from/to JSON.
type Cache struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}
