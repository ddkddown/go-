package main

import "fmt"

func test(ch <-chan int) {
	for {
		select {
		case x := <-ch:
			fmt.Println(x)
		}
	}
}

func main() {
	c := make(chan int)
	go test(c)

	for _, i := range []int{1, 2, 3} {
		c <- i
	}
}
