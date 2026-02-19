package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type webSocketHandler struct {
	upgrader websocket.Upgrader
}

func main() {
	webSocketHandler := webSocketHandler{
		upgrader: websocket.Upgrader{},
	}
	http.Handle("/", webSocketHandler)
	log.Print("Starting server...")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}

func (wsh webSocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := wsh.upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatal(err)
		return
	}

	go HandleIncomingMessages(conn)

	go HandleOutgoingMessages(conn)

}

func HandleIncomingMessages(conn *websocket.Conn) {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			continue
		}

		fmt.Printf(string(p))
	}
}

func HandleOutgoingMessages(conn *websocket.Conn) {
	for {
		if err := conn.WriteMessage(websocket.TextMessage, []byte("[server]: Testing connection!")); err != nil {
			continue
		}
		time.Sleep(1 * time.Second)
	}
}
