package main

import "net"

type room struct {
	name    string
	clients map[net.Addr]*client
}
