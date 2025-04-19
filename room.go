package main

import "net"

type room struct {
	name    string
	clients map[net.Addr]*client
}

func (r *room) broadcast(sender *client, msg string) {
	for addr, m := range r.clients {
		if addr != sender.conn.RemoteAddr() {
			m.msg(msg)
		}
	}
}
