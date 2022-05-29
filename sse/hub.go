package sse

import "bufio"

type Hub struct {
	clients map[*bufio.Writer]Client
	rooms   map[string]int

	chanRegister   chan Client
	chanBroadcast  chan Message
	chanUnregister chan *bufio.Writer
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[*bufio.Writer]Client),
		rooms:   make(map[string]int),

		chanRegister:   make(chan Client),
		chanUnregister: make(chan *bufio.Writer),
		chanBroadcast:  make(chan Message),
	}
}

func (ws Hub) GetClients() map[*bufio.Writer]Client {
	return ws.clients
}

func (ws Hub) GetClient(c *bufio.Writer) Client {
	return ws.clients[c]
}

func (ws Hub) Rooms() map[string]int {
	return ws.rooms
}

func (ws *Hub) ChanRegister() chan Client {
	return ws.chanRegister
}

func (ws *Hub) SendRegister(c Client) {
	ws.chanRegister <- c
}

func (ws *Hub) ChanUnregister() chan *bufio.Writer {
	return ws.chanUnregister
}

func (ws *Hub) SendUnregister(c *bufio.Writer) {
	ws.chanUnregister <- c
}

func (ws *Hub) ChanBroadcast() chan Message {
	return ws.chanBroadcast
}

func (ws *Hub) SendBroadcast(m Message) {
	ws.chanBroadcast <- m
}

func (ws *Hub) AddClient(c Client) {
	ws.clients[c.connection] = c
	ws.rooms[c.room]++
}

func (ws Hub) HasClient(c *bufio.Writer) bool {
	_, ok := ws.clients[c]
	return ok
}

func (ws *Hub) RemoveClient(c *bufio.Writer) {
	ws.rooms[ws.clients[c].room]--
	delete(ws.clients, c)
}

func (ws *Hub) CloseRoom(room string) {
	for c := range ws.clients {
		if ws.clients[c].room == room {
			ws.RemoveClient(c)
		}
	}

	delete(ws.rooms, room)
}
