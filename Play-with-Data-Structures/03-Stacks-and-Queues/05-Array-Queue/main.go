package main

import (
	"array"
	"fmt"
)

func main() {

	var newAS = array.NewDefaultArray()

	for i := 0; i <= 15; i++ {
		newAS.Enqueue(i)
		fmt.Println(newAS.ToString())
		if i%3 == 2 {
			newAS.Dequeue()
			fmt.Println(newAS.ToString())
		}
	}
}
