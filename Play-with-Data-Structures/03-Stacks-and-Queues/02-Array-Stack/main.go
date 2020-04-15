package main

import (
	"array"
	"fmt"
)

func main() {

	var newAS = array.NewDefaultArray()

	for i := 0; i <= 5; i++ {
		newAS.Push(i)
	}
	fmt.Println(newAS.ToString())
}
