package main

import (
	"fmt"
)

// recover 只能捕获一级内的异常错误，超出则捕获不到。

func test() {
	defer func() {
		switch p := recover(); {
		default:
			fmt.Println(p)
		}
	}()

	//test2抛出的异常在test的recover捕获范围内，所以可以捕获
	test2()
}

func test2() {
	panic("test2")
}

func main() {
	test()
}
