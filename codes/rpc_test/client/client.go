package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("dialing:", err)
	}

	var reply string
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(reply)
}
