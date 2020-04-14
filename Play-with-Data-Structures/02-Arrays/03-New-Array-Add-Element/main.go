package main

import (
	"array"
)

func main() {

	array2 := array.NewDefaultArray()
	array2.PrintAarray()

	array2.AddLast(6)
	array2.PrintAarray()

	array2.Add(0, 7)
	array2.PrintAarray()

	array2.Add(1, 8)
	array2.PrintAarray()

}
