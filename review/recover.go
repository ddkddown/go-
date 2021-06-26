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

	panic("test")
}

func test2() {
	panic("test2")
}

func main() {
	q := [3]int{0}
	w := &q
	q[0] = 2
	fmt.Println(w)
	fmt.Println(q)
	fmt.Println(w.(type))

	const i = "test"
	test()

	//捕获不了test2
	test2()
	// 此时test2内的异常已经被抛出，未捕获到
	defer func() {
		switch p := recover(); {
		default:
			fmt.Println(p)
		}
	}()

}
