package main

import (
	"array"
	"fmt"
)

func main() {

	array5 := array.NewDefaultArray()
	array5.PrintAarray()

	array5.AddLast(6)
	array5.PrintAarray()

	array5.Add(0, 7)
	array5.PrintAarray()

	array5.Add(1, 8)
	array5.PrintAarray()

	array5.Set(2, 1024)
	fmt.Printf("String list of the array elements is: %s \n", array5.ToString())
	fmt.Printf("The value of %d element in array is: %d\n", 2, array5.Get(2))
	array5.Add(3, 2048)

	array5.PrintAarray()

	array5.RemoveLast()
	array5.Add(3, 3096)
	array5.PrintAarray()

	array5.RemoveElement(3096)

	array5.PrintAarray()

}
