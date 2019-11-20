package redis

import (
	"encoding/json"
)

// UnmarshalModel unmarshals any generic struct.
func UnmarshalModel(result string) (*map[string]interface{}, error) {
	model := new(map[string]interface{})
	if err := json.Unmarshal([]byte(result), &model); err != nil {
		return nil, err
	}
	return model, nil
}

// MarshalModel marshals any generic struct or interface.
func MarshalModel(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}
