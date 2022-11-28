package websocket

import (
	"fmt"
	"log"
)

const welcomeMessage = "%s joined the room"

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
			room.registerClientInRoom(client)

		case client := <-room.unregister:
			room.unregisterClientInRoom(client)

		case message := <-room.broadcast:
			room.broadcastToClientsInRoom(message.encode())

		case message := <-room.broadcast:
			room.broadcastToClientsInRoom(message.encode())

		}
	}
}

func (room *Room) registerClientInRoom(client *Client) {
	log.Println("Registering client to room")
	room.notifyClientJoined(client)
	room.clients[client] = true
}

func (room *Room) unregisterClientInRoom(client *Client) {
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

func (room *Room) notifyClientJoined(client *Client) {
	message := &Message{
		Action:  SendMessageAction,
		Target:  room.name,
		Message: fmt.Sprintf(welcomeMessage, client.GetName()),
	}
	log.Println("notifyClientJoined message", message)
	room.broadcastToClientsInRoom(message.encode())

}
