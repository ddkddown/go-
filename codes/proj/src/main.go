package main

import (
	"ddk/myproto"
	"fmt"

	"github.com/golang/protobuf/proto"
)

func main() {
	test := &myproto.Test{
		Name: string("test"),
	}

	in_data, err := proto.Marshal(test)
	if err != nil {
		fmt.Println("Marshaling error: ", err)
	} else {
		fmt.Println(in_data)
	}

	test2 := &myproto.Test{}
	err2 := proto.Unmarshal(in_data, test2)
	if err2 != nil {
		fmt.Println("UnMarshaling error: ", err2)
	} else {
		fmt.Println(test2.Name)
	}
}
