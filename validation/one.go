package validation

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
