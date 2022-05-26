package vel

import (
	"fmt"
	"strings"
)

type ValidationErrors map[string]string

type ValidatorFunc[T any] func(T) (bool, string)

// Collection validates collection.
func Collection()

// Each run all validators on given value. It is in conjunction form.
func Each[T any](item T, callbacks ...ValidatorFunc[T]) (bool, string) {
	for _, callback := range callbacks {
		valid, err := callback(item)

		if !valid {
			return false, err
		}
	}

	return true, ""
}

// One run all validators on given value. It is in disjunction form.
func One[T any](item T, callbacks ...ValidatorFunc[T]) (bool, string) {
	var (
		valid bool
		err   string
	)

	for _, callback := range callbacks {
		valid, err = callback(item)

		if valid {
			return true, ""
		}
	}

	return false, err
}

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
