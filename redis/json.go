package redis

import (
	"encoding/json"
)

// unmarshalModel unmarshals any generic struct.
func unmarshalModel(result string) (interface{}, error) {
	model := new(interface{})
	if err := json.Unmarshal([]byte(result), &model); err != nil {
		return nil, err
	}
	return model, nil
}

// marshalModel marshals any generic struct or interface.
func marshalModel(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}
