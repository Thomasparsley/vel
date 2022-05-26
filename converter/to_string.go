package converter

import "fmt"

// ToString convert the input to a string.
func ToString(input any) string {
	return fmt.Sprintf("%v", input)
}
