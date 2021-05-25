package main 

import (
	"github.com/golang/protobuf/proto"
	"fmt"
	"os"
	"ddk/myproto"
)

func main() {
	test := &myproto.Test {
		name: proto.string("test")
	}

	in_data, err := proto.Marshal(test)
	if err != nil {
		fmt.Println("Marshaling error: ", err)
	} else {
		fmt.Println(in_data)
	}
}

