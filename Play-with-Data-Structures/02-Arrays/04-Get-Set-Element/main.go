package main

import (
	"array"
	"fmt"
)

func main() {

	array4 := array.NewDefaultArray()
	array4.PrintAarray()

	array4.AddLast(6)
	array4.PrintAarray()

	array4.Add(0, 7)
	array4.PrintAarray()

	array4.Add(1, 8)
	array4.PrintAarray()

	array4.Set(2, 1024)
	fmt.Printf("String list of the array elements is: %s \n", array4.ToString())
	fmt.Printf("The value of %d element in array is: %d\n", 2, array4.Get(2))
}
