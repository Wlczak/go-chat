package main

import (
	"bufio"
	"net"
	"strings"
)

type client struct {
	conn     net.Conn
	nick     string
	room     *room
	commands chan<- command
}

func (c *client) readInput() {

	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			return
		}

		msg = strings.Trim(msg, "\n\r")

		args := strings.Split(msg, " ")

		cmd := strings.TrimSpace(args[0])

		switch cmd {
		case "/nick":
			c.commands <- command{CMD_NICK, c, args}
		case "/join":
			c.commands <- command{CMD_JOIN, c, args}
		case "/rooms":
			c.commands <- command{CMD_ROOMS, c, args}
		case "/msg":
			c.commands <- command{CMD_MSG, c, args}
		case "/quit":
			c.commands <- command{CMD_QUIT, c, args}
		}
	}
}
