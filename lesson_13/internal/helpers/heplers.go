package helpers

import "encoding/json"

func JsonToMap(jsonString string) (map[string]any, error) {
	m := make(map[string]any)
	err := json.Unmarshal([]byte(jsonString), &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
