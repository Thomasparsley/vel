package converter

import (
	"fmt"
	"reflect"
	"strconv"
)

// ToUint64 convert the input to an uint64.
func ToUint64(input any) (uint64, error) {
	var (
		result uint64
		err    error
		value  = reflect.ValueOf(input)
	)

	switch input.(type) {
	case int, int8, int16, int32, int64:
		result = uint64(value.Int())

	case uint, uint8, uint16, uint32, uint64:
		result = value.Uint()

	case float32, float64:
		result = uint64(value.Float())

	case string:
		result, err = strconv.ParseUint(value.String(), 0, 64)
		if err != nil {
			result = 0
		}

	default:
		return 0, fmt.Errorf("ToUint: unknown interface type %T", input)
	}

	return result, nil
}

// ToUint32 convert the input to an uint32.
func ToUint32(input any) (uint32, error) {
	result, err := ToUint64(input)
	return uint32(result), err
}
