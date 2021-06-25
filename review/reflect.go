package main

import (
	"fmt"
	"reflect"
)

func main() {
	w := 1
	fmt.Println(reflect.TypeOf(w), reflect.ValueOf(w))
	v := reflect.ValueOf(w)
	x := v.Interface()
	i := x.(int)

	fmt.Println(i)
}
