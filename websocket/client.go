package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn     *websocket.Conn
	wsServer *WsServer
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
}

func NewClient(conn *websocket.Conn, wsServer *WsServer) *Client {

	client := &Client{
		conn:     conn,
		wsServer: wsServer,
	}

	return client
}

// ServeWs handles websocket requests from clients requests.
func ServeWs(wsServer *WsServer, w http.ResponseWriter, r *http.Request) {
	log.Println("server ws is called")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading websocket connection", err)
	}

	client := NewClient(conn, wsServer)

	fmt.Println("New Client joined the ws!")
	fmt.Println(client)

	wsServer.register <- client
}
