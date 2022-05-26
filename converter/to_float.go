package converter

import (
	"fmt"
	"reflect"
	"strconv"
)

// ToFloat64 convert the input to a float64.
func ToFloat64(input any) (float64, error) {
	var (
		result float64
		err    error
		value  = reflect.ValueOf(input)
	)

	switch input.(type) {
	case int, int8, int16, int32, int64:
		result = float64(value.Int())

	case uint, uint8, uint16, uint32, uint64:
		result = float64(value.Uint())

	case float32, float64:
		result = value.Float()

	case string:
		result, err = strconv.ParseFloat(value.String(), 64)
		if err != nil {
			return 0, err
		}

	default:
		return 0, fmt.Errorf("ToFloat: unknown interface type %T", input)
	}

	return result, nil
}

// ToFloat32 convert the input to a float32.
func ToFloat32(input any) (float32, error) {
	result, err := ToFloat64(input)
	return float32(result), err
}
