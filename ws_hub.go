package vel

import "github.com/gofiber/websocket/v2"

type WsHub struct {
	clients        map[*websocket.Conn]WsConnection
	registerChan   chan WsConnection
	unregisterChan chan *websocket.Conn
	broadcastChan  chan WsMessage
	rooms          map[string]int
}

func NewWsHub() *WsHub {
	return &WsHub{
		clients:        make(map[*websocket.Conn]WsConnection),
		registerChan:   make(chan WsConnection),
		unregisterChan: make(chan *websocket.Conn),
		broadcastChan:  make(chan WsMessage),
		rooms:          make(map[string]int),
	}
}

func (ws *WsHub) AddClient(c WsConnection) {
	ws.clients[c.connection] = c
	ws.rooms[c.room]++
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
