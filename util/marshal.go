package util

import "encoding/json"

func MergeAndMarshal(args ...interface{}) (string, error) {
	merged := make(map[string]interface{})
	for _, arg := range args {
		jsonBytes, err := json.Marshal(arg)
		if err != nil {
			return "", err
		}
		var m map[string]interface{}
		if err := json.Unmarshal(jsonBytes, &m); err != nil {
			return "", err
		}
		for k, v := range m {
			merged[k] = v
		}
	}
	result, err := json.Marshal(merged)
	if err != nil {
		return "", err
	}
	return string(result), nil
}
