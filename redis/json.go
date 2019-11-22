package redis

import (
	"encoding/json"

	"github.com/AquoDev/simple-imageboard-golang/model"
)

// unmarshalModel unmarshals any generic struct.
func unmarshalModel(result string) (*model.Cache, error) {
	model := new(model.Cache)
	if err := json.Unmarshal([]byte(result), &model); err != nil {
		return nil, err
	}
	return model, nil
}

// marshalModel marshals any generic struct or interface.
func marshalModel(cache *model.Cache) ([]byte, error) {
	return json.Marshal(cache)
}
