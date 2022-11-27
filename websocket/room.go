package websocket

import "log"

type Message struct {
	message []byte
}

type Room struct {
	name       string
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan *Message
}

func NewRoom(name string) *Room {
	return &Room{
		name:       name,
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan *Message),
	}
}

func (room *Room) RunRoom() {
	for {
		select {
		case client := <-room.register:
			room.registerClientToRoom(client)

		case client := <-room.unregister:
			room.unregisterClientToRoom(client)

		case message := <-room.broadcast:
			room.broadcastToClientsInRoom(message.message)

		}
	}
}

func (room *Room) registerClientToRoom(client *Client) {
	log.Println("Registering client to room")
	// room.notifyClientJoined(client)
	room.clients[client] = true
}

func (room *Room) unregisterClientToRoom(client *Client) {
	log.Println("unRegistering client to room")
	if _, ok := room.clients[client]; ok {
		delete(room.clients, client)
	}
}

func (room *Room) broadcastToClientsInRoom(message []byte) {
	log.Println("broadcast client to room")
	for client := range room.clients {
		client.send <- message
	}
}

func (room *Room) GetName() string {
	return room.name
}
