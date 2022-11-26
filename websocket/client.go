package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
}

func NewClient(conn *websocket.Conn) *Client {

	client := &Client{
		conn: conn,
	}

	return client
}

func ServeWs(w http.ResponseWriter, r *http.Request) {
	log.Println("server ws is called")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading websocket connection", err)
	}

	client := NewClient(conn)

	fmt.Println("New Client joined the ws!")
	fmt.Println(client)
}
