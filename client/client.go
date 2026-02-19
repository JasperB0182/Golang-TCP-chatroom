package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	dialer := websocket.Dialer{}
	ws, _, err := dialer.Dial("ws://localhost:8080", http.Header{})

	if err != nil {
		log.Fatal(err)
	}

	if err := ws.WriteMessage(websocket.TextMessage, []byte("hello, world!\n")); err != nil {
		log.Fatal(err)
	}

	_, p, err := ws.ReadMessage()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf(string(p))
}
