package main

import "array"

func main() {

	var newAS = array.NewDefaultArray()

	for i := 0; i <= 15; i++ {
		newAS.Enqueue(i)
		newAS.PrintAarray()
		if i%3 == 2 {
			newAS.Dequeue()
			newAS.PrintAarray()
		}
	}
}
