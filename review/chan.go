package main

import (
	"fmt"
	"time"
)

func test(ch <-chan int) {
	for {
		select {
		//没有default则会阻塞在这里
		case x := <-ch:
			fmt.Println(x)
			//default: //有default则不会阻塞，会直接执行default
			//fmt.Println("default...")
		}

		fmt.Println("outside...")
	}
}

func main() {
	c := make(chan int)
	go test(c)
	time.Sleep(time.Second * 10)
	for _, i := range []int{1, 2, 3} {
		c <- i
	}
}
