package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/common-nighthawk/go-figure"

	"portScanner/scanner"
)

func main() {

	/*
		Defining our flags

		hostname = takes a hostname value of type string

		timeout = takes a timeout value of type integer

		port scanner is extremely basic currently
	*/

	//inet := flag.String("ip", "", "an IPv4 string")
	hostname := flag.String("hostname", "", "a Hostname string")
	timeout := flag.Duration("time", 100*time.Millisecond, "Defines a timeout")
	flag.Parse()

	ips, _ := net.LookupIP(*hostname)

	executeTime := time.Now().Format(time.RFC850)

	fmt.Println(strings.Repeat("=", 75))

	myFigure := figure.NewFigure("goScanner", "", true)
	myFigure.Print()

	fmt.Println(strings.Repeat("=", 75))

	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			fmt.Printf("Starting goScan v1.0 on %s\n", executeTime)
			fmt.Printf("Checking ports on %s\n", ipv4)
		}
	}

	for port := 1; port <= 1024; port++ {

		// If condition calls on CheckPort function within the scanner package
		// CheckPort requires 3 parameters: hostname, port and timeout

		if scanner.CheckPort(*hostname, port, *timeout) {
			fmt.Printf("%d/tcp\t\topen\n", port)
		}
	}
}
