package converter

import "github.com/bytedance/sonic"

// ToJsonBytes convert the input to a valid JSON array bytes.
func ToJsonBytes(input any) ([]byte, error) {
	bytes, err := sonic.Marshal(input)
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

// FromJson conver JsonBytes to input type
func FromJson(bytes []byte, v any) error {
	return sonic.Unmarshal(bytes, v)
}
