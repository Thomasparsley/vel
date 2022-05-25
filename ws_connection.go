package vel

import "github.com/gofiber/websocket/v2"

type WsClient struct {
	connection *websocket.Conn
	sended     map[string]bool
	room       string
}

func NewWsClient(c *websocket.Conn, room string) WsClient {
	return WsClient{
		connection: c,
		sended:     make(map[string]bool),
		room:       room,
	}
}

func (c WsClient) SendString(data string) error {
	return c.connection.WriteMessage(websocket.TextMessage, []byte(data))
}
