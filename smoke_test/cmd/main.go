package main

import (
	"io"
	"log"
	"net"
)

const PORT = ":8080"

func main() {
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("error starting the tcp server: %v", err)
	}
	defer listener.Close()

	log.Printf("tcp server running on port %s", PORT)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("error receiving connection: %v", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	io.Copy(conn, conn)
}
