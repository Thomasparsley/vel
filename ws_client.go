package vel

import "github.com/gofiber/websocket/v2"

type WsClient struct {
	room       string
	connection *websocket.Conn
	sended     map[string]bool
}

func NewWsClient(c *websocket.Conn, room string) WsClient {
	return WsClient{
		room:       room,
		connection: c,
		sended:     make(map[string]bool),
	}
}

func (c WsClient) Room() string {
	return c.room
}

func (c WsClient) Sended(key string) bool {
	return c.sended[key]
}

func (c WsClient) SetSended(key string, value bool) {
	c.sended[key] = value
}

func (c WsClient) SendBytes(data []byte) error {
	return c.connection.WriteMessage(websocket.TextMessage, data)
}

func (c WsClient) SendString(data string) error {
	return c.SendBytes([]byte(data))
}
