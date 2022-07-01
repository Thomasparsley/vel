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

func None[T any]() Optional[T] {
	return Optional[T]{
		isSet: false,
	}
}

func (o Optional[T]) Value() T {
	if o.IsNone() {
		panic("cannot get value from option, is not set")
	}

	return o.value
}

func (o Optional[T]) IsSome() bool {
	return o.isSet
}

func (o Optional[T]) IsNone() bool {
	return !o.IsSome()
}
