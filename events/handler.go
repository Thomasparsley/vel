package events

type Handler[T any] func(PayloadDefiniton[T])
