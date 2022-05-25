package vel

import "github.com/gofiber/websocket/v2"

type WsConnection struct {
	connection *websocket.Conn
	sended     map[string]bool
	room       string
}

func NewWsConnection(c *websocket.Conn, room string) WsConnection {
	return WsConnection{
		connection: c,
		sended:     make(map[string]bool),
		room:       room,
	}
}
