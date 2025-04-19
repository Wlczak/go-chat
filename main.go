package main

import (
	"log"
	"net"
)

func main() {
	s := newServer()

	port := ":8888"

	listener, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Error listening: %s", err.Error())
	}

	defer listener.Close()
	log.Printf("Started server on %v", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("unable to accept connection: %s", err.Error())
			continue
		}

		go s.newClient(conn)
	}
}
