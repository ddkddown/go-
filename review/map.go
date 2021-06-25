package main

import (
	"fmt"
)

func main() {

	m := map[string]int{
		"test":  1,
		"test2": 2,
		"test3": 3,
		"test4": 4,
		"test5": 5,
	}

	fmt.Println(len(m))

	for k, v := range m {
		fmt.Println(k, v)
	}

	m["test6"] = 6
	for k, v := range m {
		fmt.Println(k, v)
	}

}
