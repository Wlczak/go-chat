package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

var serial_mu sync.Mutex

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
	fmt.Print("\033[0G")
	fmt.Print("\033[K")
	fmt.Print("> ")
	serial_mu.Unlock()
	if scanner.Scan() {
		text := scanner.Text()
		fmt.Print("\033[A\r\033[K")
		return text
	} else {
		return ""
	}
}

func connect() (net.Conn, error) {
	ip := "localhost:8888"
	conn, err := net.Dial("tcp", ip)

	if err != nil {
		fmt.Println(err)
		return conn, err
	}
	return conn, nil
}
