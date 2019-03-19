package main

import (
	"fmt"
)

func A() int {
	a := 5
	defer func(b *int) {
		*b += 1
	}(&a)

	return a
}

func main() {
	fmt.Println(A())
	fmt.Println(B())

}
func B() int {
	a := 3
	defer func() {
		a += 1
	}()
	return a

}
