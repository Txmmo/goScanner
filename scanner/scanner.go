package scanner

import (
	"fmt"
	"net"
	"time"
)

func CheckPort(host string, port int, timeout time.Duration ) bool {
	address := fmt.Sprintf("%s:%d", host, port)

	conn, err := net.DialTimeout("tcp", address, timeout)

	if err != nil {
		return false
	}

	conn.Close()
	
	return true
}