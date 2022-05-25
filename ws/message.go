package ws

type Message struct {
	data  string
	room  string
	extra map[string]bool
}

func NewMessage(data string, room string) Message {
	return Message{
		data:  data,
		room:  room,
		extra: make(map[string]bool),
	}
}

func (m Message) Data() string {
	return m.data
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
