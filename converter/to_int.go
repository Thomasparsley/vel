package converter

import (
	"fmt"
	"reflect"
	"strconv"
)

// ToInt64 convert the input to an int64.
func ToInt64(input any) (int64, error) {
	var (
		result int64
		err    error
		value  = reflect.ValueOf(input)
	)

	switch input.(type) {
	case int, int8, int16, int32, int64:
		result = value.Int()

	case uint, uint8, uint16, uint32, uint64:
		result = int64(value.Uint())

	case float32, float64:
		result = int64(value.Float())

	case string:
		result, err = strconv.ParseInt(value.String(), 0, 64)
		if err != nil {
			result = 0
		}

	default:
		return 0, fmt.Errorf("ToInt: unknown interface type %T", input)
	}

	return result, nil
}

// ToInt32 convert the input to an int32.
func ToInt32(input any) (int32, error) {
	result, err := ToInt64(input)
	return int32(result), err
}
