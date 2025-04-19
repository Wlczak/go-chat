package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

var serial_mu sync.Mutex

const (
	msg_prefix = "> "
)

func main() {
	conn, err := connect()

	if err != nil {
		serial_mu.Lock()
		fmt.Print(err)
		serial_mu.Unlock()
		return
	}

	go readConn(conn)

	for {
		send_msg := getInput()
		conn.Write([]byte(send_msg + "\n"))
	}
}

func readConn(conn net.Conn) {
	conn_reader := bufio.NewReader(conn)
	for {
		recieve_msg, err := conn_reader.ReadString('\n')

		if err != nil {
			print(fmt.Sprintf("%v", err))
			return
		}
		print(fmt.Sprint(recieve_msg))
	}

}

func getInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	serial_mu.Lock()
	fmt.Print(msg_prefix)
	serial_mu.Unlock()
	if scanner.Scan() {
		text := scanner.Text()
		rmLine()
		return filterInput(text)
	} else {
		return ""
	}
}

func filterInput(msg string) string {
	args := strings.Split(msg, " ")

	cmd := args[0]

	switch cmd {
	case "/nick":
		return "/nick " + args[1]
	case "/join":
		return "/join " + args[1]
	case "/rooms":
		return "/rooms"
	case "/msg":
		return "/msg " + strings.Join(args[1:], " ")
	case "/quit":
		return "/quit"
	default:
		return "/msg " + msg
	}

	return msg
}

func connect() (net.Conn, error) {
	ip := "192.168.0.105:8888"
	conn, err := net.Dial("tcp", ip)

	if err != nil {
		fmt.Println(err)
		return conn, err
	}
	return conn, nil
}
