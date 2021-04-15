package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("ListenTCP error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Accept error:", err)
	}

	rpc.ServeConn(conn)
}
