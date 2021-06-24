package main

import "fmt"

func test(x interface{}) {
	if x == nil {
		fmt.Println("nil")
	} else if t, ok := x.(uint); ok {
		fmt.Println("uint:", t)
	} else if t, ok := x.(int); ok {
		fmt.Println("int:", t)
	}

}
func main() {
	var x int
	x = 1
	test(x)
}
