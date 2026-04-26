package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	
	timeout := 1*time.Second

	for port := 1; port <= 1024; port++ {

		address := fmt.Sprintf("scanme.nmap.org:%d", port)

		conn, err := net.DialTimeout("tcp", address, timeout)

		if err != nil {
			continue // port closed/filtered → move on
		}

		conn.Close()
		fmt.Printf("Port %d is open\n", port)
	}
}