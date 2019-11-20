package redis

import (
	"encoding/json"

	"github.com/AquoDev/simple-imageboard-golang/model"
)

// UnmarshalModel unmarshals any generic struct.
func UnmarshalModel(result string, model interface{}) (err error) {
	cachedModel := []byte(result)
	err = json.Unmarshal(cachedModel, &model)
	return
}

// UnmarshalPostSlice unmarshals a post slice.
func UnmarshalPostSlice(result string) ([]model.Post, error) {
	postSlice := make([]model.Post, 0)
	if err := UnmarshalModel(result, &postSlice); err != nil {
		return nil, err
	}
	return postSlice, nil
}

// UnmarshalPost unmarshals a single post.
func UnmarshalPost(result string) (*model.Post, error) {
	post := new(model.Post)
	if err := UnmarshalModel(result, &post); err != nil {
		return nil, err
	}
	return post, nil
}

// MarshalModel marshals any generic struct or interface.
func MarshalModel(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}
