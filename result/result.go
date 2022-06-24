package result

type Result[T any] struct {
	value T
	err   error
}

func (r Result[T]) IsOk() bool {
	return r.err == nil
}

func (r Result[T]) Value() T {
	if !r.IsOk() {
		panic("cannot get value from result, error is present")
	}

	return r.value
}

func (r Result[T]) Error() error {
	return r.err
}
