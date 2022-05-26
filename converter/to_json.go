package converter

import "encoding/json"

// ToJsonBytes convert the input to a valid JSON array bytes.
func ToJsonBytes(input any) ([]byte, error) {
	bytes, err := json.Marshal(input)
	if err != nil {
		return []byte{}, err
	}

	return bytes, err
}

// ToJson convert the input to a valid JSON string.
func ToJson(input any) (string, error) {
	bytes, err := ToJsonBytes(input)
	if err != nil {
		return "", err
	}

	return string(bytes), err
}
