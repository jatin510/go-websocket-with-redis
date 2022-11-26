package main

import (
	"flag"
	"log"
	"net/http"
	"websocket-server/websocket"
)

var addr = flag.String("addr", ":8090", "http server address")

func main() {
	flag.Parse()

	wsServer := websocket.NewWebsocketServer()
	go wsServer.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket.ServeWs(wsServer, w, r)
	})

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	log.Println("server started")
	log.Fatal(http.ListenAndServe(*addr, nil))
}
