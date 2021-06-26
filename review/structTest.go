package main

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

type t struct{}

func main() {

	t1 := t{}
	fmt.Println(unsafe.Sizeof(t1))

	t2, _ := json.Marshal(t1)
	fmt.Println(string(t2))

	/*
		a := structer.Test{1, 1} //编译报无法引用错误
		b := structer.Test2{1, 1}
	*/
}
