package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	test := func() {
		fmt.Println("only once")
	}

	once.Do(test)
	once.Do(test)
	once.Do(test)
}
