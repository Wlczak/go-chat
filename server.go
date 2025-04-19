package main

import (
	"log"
	"net"
)

type server struct {
	rooms    map[string]*room
	commands chan command
}

func newServer() *server {

	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),
	}
}

func (s *server) newClient(conn net.Conn) {
	log.Printf("New client: %s", conn.RemoteAddr().String())

	c := &client{
		conn:     conn,
		nick:     "anonymous",
		commands: s.commands,
	}

	c.readInput()
}
