package main

import "fmt"

type fuck interface {
	call()
}

type fucker struct {
}

func (person *fucker) call() {
	fmt.Println("fuck you")
}

func main() {
	var test fuck
	test = new(fucker)
	test.call()
}
