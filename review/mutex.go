package main

import (
	"fmt"
	"sync"
)

var (
	mu sync.Mutex
	wg sync.WaitGroup
)

func test1() {
	mu.Lock()
	defer mu.Unlock()
	defer wg.Done()
	fmt.Println("test1")
}

func test2() {
	mu.Lock()
	defer mu.Unlock()
	defer wg.Done()
	fmt.Println("test2")
}

func main() {
	wg.Add(2)

	go test1()
	go test2()

	wg.Wait()
}
