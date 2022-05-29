package events

type PayloadDefiniton[T any] interface {
	GetPayload() T
}
