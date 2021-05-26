package main

import (
	"context"
	"fmt"
	"log"

	"ddk/myproto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := myproto.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &myproto.Test{Name: "ddk"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply.GetName())
}
