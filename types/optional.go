package types

import (
	"database/sql/driver"
	"errors"
	"fmt"

	"github.com/Thomasparsley/vel/converter"
)

type OptionalState bool

const (
	OptionalState_Set   = OptionalState(true)
	OptionalState_Unset = OptionalState(false)
)

func errUnmarshalValue(value any) error {
	return errors.New(
		fmt.Sprint("Failed to unmarshal Optional value:", value),
	)
}

var (
	ErrUnmarshalValue = errUnmarshalValue
)

type Optional[T any] struct {
	value T
	state OptionalState
}

func Some[T any](value T) Optional[T] {
	return Optional[T]{
		value: value,
		state: OptionalState_Set,
	}
}

func None[T any]() Optional[T] {
	return Optional[T]{
		state: OptionalState_Unset,
	}
}

func (o Optional[T]) Unwrap() T {
	if o.IsNone() {
		panic("cannot get value from option, is not set")
	}

	return o.value
}

func (o Optional[T]) IsSome() bool {
	return o.state == OptionalState_Set
}

func (o Optional[T]) IsNone() bool {
	return !o.IsSome()
}

// Value implements the driver Valuer interface.
func (o Optional[T]) Value() (driver.Value, error) {
	if o.IsNone() {
		return nil, nil
	}

	return o.value, nil
}

// Scan implements the Scanner interface.
func (o *Optional[T]) Scan(src any) error {
	if src == nil {
		o.state = OptionalState_Unset
		return nil
	}

	value, ok := src.(T)
	if !ok {
		return ErrUnmarshalValue(src)
	}

	o.value = value
	o.state = OptionalState_Set
	return nil
}

func (o Optional[T]) MarshalJSON() ([]byte, error) {
	if o.IsNone() {
		return converter.ToJsonBytes(o.value)
	}

	return converter.ToJsonBytes(nil)
}

func (o *Optional[T]) UnmarshalJSON(src []byte) error {
	o.state = OptionalState_Unset

	if string(src) == "null" {
		return nil
	}

	err := converter.FromJson(src, &o.value)
	if err == nil {
		o.state = OptionalState_Set
	}

	return err
}
