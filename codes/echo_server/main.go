package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	defer c.Close()
	echo(c, "test", 1*time.Second)
}

func main() {
	lisener, _ := net.Listen("tcp", "localhost:8000")
	for {
		conn, _ := lisener.Accept()
		handleConn(conn)
	}
}
