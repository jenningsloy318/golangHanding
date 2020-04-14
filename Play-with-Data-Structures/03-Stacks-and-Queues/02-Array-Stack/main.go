package main

import "array"

func main() {

	var newAS = array.NewDefaultArray()

	for i := 0; i <= 5; i++ {
		newAS.Push(i)
	}
	newAS.PrintAarray()
}
