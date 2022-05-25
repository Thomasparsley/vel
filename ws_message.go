package vel

type WsMessage struct {
	data  string
	room  string
	extra map[string]bool
}

func NewWsMessage(data string, room string) WsMessage {
	return WsMessage{
		data:  data,
		room:  room,
		extra: make(map[string]bool),
	}
}

func (m WsMessage) Data() string {
	return m.data
}

func (m WsMessage) Room() string {
	return m.room
}

func (m *WsMessage) SetRoom(room string) {
	m.room = room
}

func (m WsMessage) Extra() map[string]bool {
	return m.extra
}

func (m *WsMessage) SetExtra(key string, value bool) {
	m.extra[key] = value
}
