package main

import (
	"log"
	"net"
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
	}
}
