package main

import (
	"fmt"
	"os"
)

func main() {
	if 1 >= len(os.Args) {
		fmt.Println("plz input args")
		os.Exit(1)
	}

	for index, arg := range os.Args {
		fmt.Printf("arg %d %s\n", index, arg)
	}
}
