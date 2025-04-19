package main

import "fmt"

func print(msg string) {
	serial_mu.Lock()
	defer serial_mu.Unlock()
	fmt.Print("\033[A")
	// Optionally, clear the line
	// fmt.Print("\033[K")
	// Move to the end of the line
	fmt.Print("\033[999C")
	fmt.Print(msg)
}
