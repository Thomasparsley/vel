package result

type State bool

const (
	STATE_SUCCESS = State(true)
	STATE_FAULTED = State(false)
)

type Result[T any] struct {
	state State
	value T
	error error
}

func Ok[T any](value T) Result[T] {
	return Result[T]{
		state: STATE_SUCCESS,
		value: value,
	}
}

func Error[T any](error error) Result[T] {
	return Result[T]{
		state: STATE_FAULTED,
		error: error,
	}
}

func Match[T any, V any](
	result Result[T],
	ok func(value T) V,
	err func(error error) V,
) V {
	if result.IsValid() {
		return ok(result.Unwrap())
	}

	return err(result.Error())
}

func (r Result[T]) IsValid() bool {
	return r.state == STATE_SUCCESS
}

func (r Result[T]) IsInvalid() bool {
	return r.state == STATE_FAULTED
}

func (r Result[T]) Unwrap() T {
	if r.IsValid() {
		return r.value
	}

	panic("cannot get value from result, value is not present")
}

func (r Result[T]) ValueOr(v T) T {
	if r.IsValid() {
		return r.value
	}

	return v
}

func (r Result[T]) Error() error {
	if r.IsInvalid() {
		return r.error
	}

	panic("cannot get error from result, error is not present")
}
