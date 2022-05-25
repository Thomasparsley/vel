package ws

import "github.com/gofiber/websocket/v2"

type Hub struct {
	clients map[*websocket.Conn]Client
	rooms   map[string]int

	chanRegister   chan Client
	chanBroadcast  chan Message
	chanUnregister chan *websocket.Conn
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[*websocket.Conn]Client),
		rooms:   make(map[string]int),

		chanRegister:   make(chan Client),
		chanUnregister: make(chan *websocket.Conn),
		chanBroadcast:  make(chan Message),
	}
}

func (ws Hub) GetClients() map[*websocket.Conn]Client {
	return ws.clients
}

func (ws Hub) GetClient(c *websocket.Conn) Client {
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

func (ws *Hub) ChanUnregister() chan *websocket.Conn {
	return ws.chanUnregister
}

func (ws *Hub) SendUnregister(c *websocket.Conn) {
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

func (ws Hub) HasClient(c *websocket.Conn) bool {
	_, ok := ws.clients[c]
	return ok
}

func (ws *Hub) RemoveClient(c *websocket.Conn) {
	c.WriteMessage(websocket.CloseMessage, []byte{})
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
