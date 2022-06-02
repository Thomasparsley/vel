package validation

import (
	"fmt"
	"strings"
)

// MaxLengthString validates string length.
func MaxLengthString(length int) ValidatorFunc[string] {
	return func(value string) (bool, string) {
		if len(value) > length {
			return false,
				fmt.Sprintf("The length of the string is too long, max length is %d", length)
		}

		return true, ""
	}
}

// MinLengthString validates string length.
func MinLengthString(length int) ValidatorFunc[string] {
	return func(value string) (bool, string) {
		if len(value) < length {
			return false,
				fmt.Sprintf("The length of the string is too short, min length is %d", length)
		}

		return true, ""
	}
}

// RequiredString validates string length.
func RequiredString() ValidatorFunc[string] {
	return func(value string) (bool, string) {
		if strings.TrimSpace(value) == "" {
			return false, "This field is required"
		}

		return true, ""
	}
}

// !
// IsEmail validates email.
func IsEmail() ValidatorFunc[string] {
	return func(s string) (bool, string) {
		if !strings.Contains(s, "@") {
			return false, "The email is not valid"
		}

		return true, ""
	}
}
