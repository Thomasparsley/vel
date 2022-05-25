package ws

import "encoding/json"

type Message struct {
	data  []byte
	room  string
	extra map[string]bool
}

func NewMessage(data []byte, room string) Message {
	return Message{
		data:  data,
		room:  room,
		extra: make(map[string]bool),
	}
}

func NewStringMessage(data string, room string) Message {
	return NewMessage([]byte(data), room)
}

func NewJsonMessage(data any, room string) Message {
	jsonBytes, _ := json.Marshal(data)

	return NewMessage(jsonBytes, room)
}

func (m Message) Data() []byte {
	return m.data
}

func (m Message) DataString() string {
	return string(m.data)
}

func (m Message) Room() string {
	return m.room
}

func (m *Message) SetRoom(room string) {
	m.room = room
}

func (m Message) Extra(key string) bool {
	return m.extra[key]
}

func (m *Message) SetExtra(key string, value bool) {
	m.extra[key] = value
}
