package main

import (
	"array"
)

func main() {

	array9 := array.NewDefaultArray()

	cap := array9.GetCapacity()

	for i := 0; i < cap; i++ {
		array9.AddLast(2 * i)
	}

	array9.PrintAarray()

	for i := 0; i < 6; i++ {
		array9.RemoveLast()
		array9.PrintAarray()
	}
}
