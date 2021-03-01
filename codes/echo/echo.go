package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string

	//1 version
	/*for i := 1; i < len(os.Args); i++ {
		s += os.Args[i]
		s += " "
	}*/
	
	//2 version
	/*
	for _, arg := range os.Args[1:] {
		s += arg + " "
	}*/

	//3 version
	s += strings.Join(os.Args[1:], " ")

	fmt.Println(s)
}
