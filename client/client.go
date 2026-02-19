package main

import (
	"bufio"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	dialer := websocket.Dialer{}
	ws, _, err := dialer.Dial("ws://localhost:8080", http.Header{})
	defer ws.Close()

	if err != nil {
		log.Fatal(err)
	}

	go OutgoingMessages(ws)
	go IncomingMessages(ws)
	select {}

}

func OutgoingMessages(ws *websocket.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')

		if err := ws.WriteMessage(websocket.TextMessage, []byte(line)); err != nil {
			log.Fatal(err)
		}
	}
}

func IncomingMessages(ws *websocket.Conn) {
	for {
		_, p, err := ws.ReadMessage()

		if err != nil {
			log.Fatal(err)
		}

		log.Printf(string(p))
	}
}
