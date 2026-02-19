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

	if err != nil {
		log.Fatal(err)
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')

		if err := ws.WriteMessage(websocket.TextMessage, []byte(line)); err != nil {
			log.Fatal(err)
		}

		_, p, err := ws.ReadMessage()

		if err != nil {
			log.Fatal(err)
		}

		log.Printf(string(p))
	}

}
