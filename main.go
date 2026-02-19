package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	const address = ":8090"
	const networkProtocol = "tcp"
	// This is legacy code since I have decided to switch to websockets for this project.

	l, err := net.Listen(networkProtocol, address)
	if err != nil {
		log.Println(err)
	}
	log.Println("Listening on ", l.Addr())

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Error: %s while connecting to addr: %s\n",
				err.Error(), conn.RemoteAddr())
		} else {
			log.Printf("Connected successfully to remote addr: %s\n",
				conn.RemoteAddr())
		}

		go HandleOutgoing(conn)

		go HandleIncoming(conn)

	}

}

func HandleOutgoing(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Could not send message")
			return
		}

		newContext := fmt.Sprintf("[Server]: %s", line)

		buffer := []byte(newContext)

		_, err = conn.Write(buffer)

		if err != nil {
			fmt.Println("Error writing to client")
			return
		}

		fmt.Printf("[Server]: %s", line)
	}

}

func HandleIncoming(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 64)

		n, err := conn.Read(buf)

		if err != nil {
			//TODO: Connection shuts down ungracefully
			log.Fatal("Connection closed.")
			return
		}

		fmt.Printf("[Client]: %s", string(buf[:n]))
	}
}
