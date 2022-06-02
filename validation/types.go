package validation

type ValidationErrors map[string]string

type ValidatorFunc[T any] func(T) (bool, string)
