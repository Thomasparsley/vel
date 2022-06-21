package validation

type Errors map[string]string

type ValidatorFunc[T any] func(T) (bool, string)
