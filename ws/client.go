package ws

import "github.com/gofiber/websocket/v2"

type Client struct {
	room       string
	connection *websocket.Conn
	sended     map[string]bool
}

func NewClient(c *websocket.Conn, room string) Client {
	return Client{
		room:       room,
		connection: c,
		sended:     make(map[string]bool),
	}
}

func (c Client) Room() string {
	return c.room
}

func (c Client) Sended(key string) bool {
	return c.sended[key]
}

func (c Client) SetSended(key string, value bool) {
	c.sended[key] = value
}

func (c Client) SendBytes(data []byte) error {
	return c.connection.WriteMessage(websocket.TextMessage, data)
}

func (c Client) SendMessage(message Message) error {
	return c.SendBytes(message.Data())
}

func (c Client) SendString(data string) error {
	return c.SendBytes([]byte(data))
}

func (c Client) SendJSON(json any) error {
	return c.connection.WriteJSON(json)
}
