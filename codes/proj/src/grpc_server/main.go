package main

import (
	"log"
	"net"

	"ddk/myproto"

	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	myproto.RegisterHelloServiceServer(grpcServer, new(HelloServiceImp))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer.Serve(lis)
}
