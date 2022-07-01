package result

import "github.com/Thomasparsley/vel/structs/optional"

type State bool

const (
	STATE_SUCCESS = State(true)
	STATE_FAULTED = State(false)
)

type Result[T any, E any] struct {
	state State
	value T
	error E
}

func Ok[T any, E any](value T) Result[T, E] {
	return Result[T, E]{
		state: STATE_SUCCESS,
		value: value,
	}
}

func Error[T any, E any](error E) Result[T, E] {
	return Result[T, E]{
		state: STATE_FAULTED,
		error: error,
	}
}

func Match[T any, E any, TResult any, EResult any](
	result Result[T, E],
	ok func(value T) TResult,
	err func(error E) EResult,
) (optional.Optional[TResult], optional.Optional[EResult]) {
	if result.IsValid() {
		return optional.Some(ok(result.Value())), optional.None[EResult]()
	}

	return optional.None[TResult](), optional.Some(err(result.Error()))
}

func (r Result[T, E]) IsValid() bool {
	return r.state == STATE_SUCCESS
}

func (r Result[T, E]) IsInvalid() bool {
	return r.state == STATE_FAULTED
}

func (r Result[T, E]) Value() T {
	if !r.IsValid() {
		panic("cannot get value from result, value is not present")
	}

	return r.value
}

func (r Result[T, E]) Error() E {
	if !r.IsInvalid() {
		panic("cannot get error from result, error is not present")
	}

	return r.error
}
