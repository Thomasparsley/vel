package optional

type Optional[T any] struct {
	value T
	isSet bool
}

func Some[T any](value T) Optional[T] {
	return Optional[T]{
		value: value,
		isSet: true,
	}
}

func None() Optional[any] {
	return Optional[any]{
		isSet: false,
	}
}

func (o Optional[T]) Value() T {
	return o.value
}

func (o Optional[T]) IsSet() bool {
	return o.isSet
}

func IsSome[T any](o Optional[T]) bool {
	return o.isSet
}

func IsNone[T any](o Optional[T]) bool {
	return !IsSome(o)
}
