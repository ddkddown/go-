package main

import "fmt"

func main() {
	t := make([]int, 3, 10)

	fmt.Println(len(t), cap(t))

	for i := 0; i < len(t); i++ {
		fmt.Println("len:", t[i])
	}

	for _, j := range t {
		fmt.Println("t elem: ", j)
	}

	for i := range []int{1, 2, 3, 4, 5} {
		fmt.Println("appending: ", i)
		t = append(t, i)
		fmt.Println(len(t), cap(t))
	}
}
