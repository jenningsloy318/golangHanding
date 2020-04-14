package main

import (
	"fmt"
)

func main() {
	var a [3]int
	for i := 0; i <= 2; i++ {
		a[i] = i
	}
	fmt.Printf("array a is %d\n", a)

	var b = [3]int{1, 2, 3}
	for i := 0; i < 3; i++ {
		fmt.Printf("b[%d] is %d\n", i, b[i])
	}

	for index, value := range b {

		fmt.Printf("index is %d, and value is %d\n", index, value)

	}
}
