package events

type Event[T PayloadDefiniton[T]] struct {
	name     string
	handlers []Handler[T]
}

func New[T PayloadDefiniton[T]](name string) Event[T] {
	return Event[T]{
		name:     name,
		handlers: make([]Handler[T], 0),
	}
}

func (e Event[T]) Name() string {
	return e.name
}

func (e *Event[T]) Register(handler Handler[T]) {
	e.handlers = append(e.handlers, handler)
}

func (e Event[T]) Trigger(payload PayloadDefiniton[T]) {
	for _, handler := range e.handlers {
		go handler(payload)
	}
}
