package main

import (
	"fmt"
	"time"

	"portScanner/scanner"
)

func main() {
	
	timeout := 1*time.Second

	for port := 1; port <= 1024; port++ {

	if scanner.CheckPort("scanme.nmap.org", port, timeout) {
		fmt.Printf("Port %d is open\n", port)
		}
	}
}