package sse

import (
	"bufio"
	"fmt"

	"github.com/Thomasparsley/vel/converter"
)

type Client struct {
	room       string
	connection *bufio.Writer
	sended     map[string]bool
}

func NewClient(c *bufio.Writer, room string) Client {
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
	fmt.Fprintf(c.connection, "data: %s\n\n", string(data))
	return c.connection.Flush()
}

func (c Client) SendMessage(message Message) error {
	return c.SendBytes(message.Data())
}

func (c Client) SendString(data string) error {
	return c.SendBytes([]byte(data))
}

func (c Client) SendJSON(json any) error {
	_json, _ := converter.ToJsonBytes(json)
	return c.SendBytes(_json)
}
