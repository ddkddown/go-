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

	fmt.Println("ready to go")
}
