package main

import (
	"array"
)

func main() {

	array7 := array.NewDefaultArray()

	cap := array7.GetCapacity()

	for i := 0; i < cap; i++ {
		array7.AddLast(2 * i)
	}

	array7.PrintAarray()

	for i := 0; i < 6; i++ {
		array7.RemoveLast()
		array7.PrintAarray()
	}
}
