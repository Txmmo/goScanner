package main

import (
	"fmt"
	"time"
	"flag"
	"strings"
	"github.com/common-nighthawk/go-figure"

	"portScanner/scanner"
)

func main() {

	/*
		Defining our flags

		hostname = takes a hostname value of type string

		timeout = takes a timeout value of type integer
	*/

	hostname := flag.String("hostname", "", "a string" )
	timeout := flag.Duration("time", 100*time.Millisecond, "Defines a timeout")
	flag.Parse()

	executeTime := time.Now().Format(time.RFC850)

	fmt.Println(strings.Repeat("=", 75))

	myFigure := figure.NewFigure("goScanner", "", true)
	myFigure.Print()

	fmt.Println(strings.Repeat("=", 75))
	fmt.Printf("Starting goScan v1.0 on %s\n", executeTime)
	fmt.Printf("Checking ports on %s\n", *hostname)

	for port := 1; port <= 1024; port++ {
	
	// If condition calls on CheckPort function within the scanner package
	// CheckPort requires 3 parameters: hostname, port and timeout

	if scanner.CheckPort(*hostname, port, *timeout) {
		fmt.Printf("%d/tcp		open\n", port)
		}
	}
}