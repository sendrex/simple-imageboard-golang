package redis

import (
	"encoding/json"

	"github.com/AquoDev/simple-imageboard-golang/model"
)

// unmarshalCache turns any JSON into a Cache struct.
func unmarshalCache(result string) (*model.Cache, error) {
	model := new(model.Cache)
	if err := json.Unmarshal([]byte(result), &model); err != nil {
		return nil, err
	}
	return model, nil
}

// marshalCache turns any Cache struct into JSON.
func marshalCache(cache *model.Cache) ([]byte, error) {
	return json.Marshal(cache)
}
