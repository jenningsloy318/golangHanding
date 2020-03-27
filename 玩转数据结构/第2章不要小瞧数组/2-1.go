package main

import "fmt"

func main() {
	var a [3]int
	for i := 0; i <= 2; i++ {
		a[i] = i
	}
	fmt.Println(a)

	var b = [3]int{1, 2, 3}
	for i := 0; i < 3; i++ {
		fmt.Println(b[i])
	}

	for index, value := range b {
		fmt.Printf("index is %d, and value is %d\n", index, value)
	}
}
