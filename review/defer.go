package main

import "fmt"

func test(i int) func() {
	return func() {
		fmt.Println(i)
	}
}

func main() {
	t := [3]int{1, 2, 3}

	for _, v := range t {
		defer test(v)()
	}

	//虽然抛出了异常，但是仍然会执行defer
	panic("test")

	fmt.Println("ready to go")
}
