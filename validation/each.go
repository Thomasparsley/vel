package validation

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
