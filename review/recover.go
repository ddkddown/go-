package main

import (
	"fmt"
)

func test() {
	defer func() {
		switch p := recover(); {
		default:
			fmt.Println(p)
		}
	}()

	panic("test")
}

func test2() {
	panic("test2")
}

func main() {
	test()

	//捕获不了test2
	test2()
	defer func() {
		switch p := recover(); {
		default:
			fmt.Println(p)
		}
	}()

}
