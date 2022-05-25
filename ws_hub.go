package vel

import "github.com/gofiber/websocket/v2"

type WsHub struct {
	clients map[*websocket.Conn]WsClient
	rooms   map[string]int

	chanRegister   chan WsClient
	chanBroadcast  chan WsMessage
	chanUnregister chan *websocket.Conn
}

func NewWsHub() *WsHub {
	return &WsHub{
		clients: make(map[*websocket.Conn]WsClient),
		rooms:   make(map[string]int),

		chanRegister:   make(chan WsClient),
		chanUnregister: make(chan *websocket.Conn),
		chanBroadcast:  make(chan WsMessage),
	}
}

func (ws WsHub) GetClients() map[*websocket.Conn]WsClient {
	return ws.clients
}

func (ws WsHub) GetClient(c *websocket.Conn) WsClient {
	return ws.clients[c]
}

func (ws WsHub) Rooms() map[string]int {
	return ws.rooms
}

func (ws *WsHub) ChanRegister() chan WsClient {
	return ws.chanRegister
}

func (ws *WsHub) SendRegister(c WsClient) {
	ws.chanRegister <- c
}

func (ws *WsHub) ChanUnregister() chan *websocket.Conn {
	return ws.chanUnregister
}

func (ws *WsHub) SendUnregister(c *websocket.Conn) {
	ws.chanUnregister <- c
}

func (ws *WsHub) ChanBroadcast() chan WsMessage {
	return ws.chanBroadcast
}

func (ws *WsHub) SendBroadcast(m WsMessage) {
	ws.chanBroadcast <- m
}

func (ws *WsHub) AddClient(c WsClient) {
	ws.clients[c.connection] = c
	ws.rooms[c.room]++
}

func (ws WsHub) HasClient(c *websocket.Conn) bool {
	_, ok := ws.clients[c]
	return ok
}

func (ws *WsHub) RemoveClient(c *websocket.Conn) {
	c.WriteMessage(websocket.CloseMessage, []byte{})
	ws.rooms[ws.clients[c].room]--
	delete(ws.clients, c)
}

func (ws *WsHub) CloseRoom(room string) {
	for c := range ws.clients {
		if ws.clients[c].room == room {
			ws.RemoveClient(c)
		}
	}

	delete(ws.rooms, room)
}
