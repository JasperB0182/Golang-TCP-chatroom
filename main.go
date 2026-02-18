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

	}

}

func HandleOutgoing(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Could not send message")
		}

		buffer := []byte(line)

		_, err = conn.Write(buffer)

		if err != nil {
			fmt.Println("Error writing to client")
		}

		fmt.Printf("[Server]: %s", line)
	}

}
