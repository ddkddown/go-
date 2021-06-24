package main

import "fmt"

//函数squares返回的是一个闭包。闭包
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

//对squares的一次调用会生成一个所对应的局部变量x
func main() {
	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"
}
