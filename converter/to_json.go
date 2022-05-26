package converter

import "encoding/json"

// ToJSON convert the input to a valid JSON string.
func ToJSON(input any) (string, error) {
	bytes, err := json.Marshal(input)
	if err != nil {
		return "", err
	}

	return string(bytes), err
}
