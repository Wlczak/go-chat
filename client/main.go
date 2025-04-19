package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	connect()
}

func getInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter text: ")
	if scanner.Scan() {
		text := scanner.Text()
		return text
	} else {
		return ""
	}
}

func connect() {
	ip := "localhost:8888"
	_, err := net.Dial("tcp", ip)

	if err != nil {
		fmt.Println(err)
		return
	}

}
