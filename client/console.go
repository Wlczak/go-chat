package main

import (
	"fmt"
	"strings"
)

func print(msg string) {
	serial_mu.Lock()
	defer serial_mu.Unlock()
	fmt.Print("\033[F")
	fmt.Print(strings.Trim(msg, "\n\r"))
	fmt.Print("\033[B")
	if !is_last {
		fmt.Printf("\033[%dD", len(msg)-len(msg_prefix)-1)
	} else {
		fmt.Printf("\033[%dD", 999)
	}

}

func rmLine() {
	serial_mu.Lock()
	fmt.Print("\033[A")
	fmt.Print("\033[K")
	fmt.Print("\033[B")
	serial_mu.Unlock()
}
