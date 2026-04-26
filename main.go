package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	
	timeout := 1*time.Second
	conn, err := net.DialTimeout("tcp", "scanme.nmap.org:80", timeout)

	if err != nil {
		fmt.Println("Port 80 is closed or filtered")
		return
	}
	defer conn.Close()
	fmt.Println("Port 80 is open")
}