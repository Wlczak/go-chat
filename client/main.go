package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

var serial_mu sync.Mutex
var wg sync.WaitGroup

var is_last bool = false

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

	welcome()

	wg.Add(1)
	go readConn(conn)

	go handleInput(conn)

	wg.Wait()
}

func welcome() {

	fmt.Println(`
 _    _      _                          
| |  | |    | |                         
| |  | | ___| | ___ ___  _ __ ___   ___ 
| |/\| |/ _ \ |/ __/ _ \| '_ ' _ \ / _ \
\  /\  /  __/ | (_| (_) | | | | | |  __/
 \/  \/ \___|_|\___\___/|_| |_| |_|\___|	
	`)

}

func handleInput(conn net.Conn) {
	for {
		msg := getInput()
		conn.Write([]byte(msg + "\n"))
	}
}

func readConn(conn net.Conn) {
	defer wg.Done()

	conn_reader := bufio.NewReader(conn)
	for {
		recieve_msg, err := conn_reader.ReadString('\n')

		if err != nil {
			//print(fmt.Sprintf("%v", err))
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
		if len(args) < 2 {
			return "/nick " + rand.Text()
		}
		return "/nick " + args[1]
	case "/join":
		if len(args) < 2 {
			return "/join " + rand.Text()
		}
		return "/join " + args[1]
	case "/rooms":
		return "/rooms"
	case "/msg":
		return "/msg " + strings.Join(args[1:], " ")
	case "/quit":
		is_last = true
		return "/quit"
	default:
		return "/msg " + msg
	}
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
