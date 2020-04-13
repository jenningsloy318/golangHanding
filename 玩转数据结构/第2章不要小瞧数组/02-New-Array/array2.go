package main

import "fmt"

//Array declare new Array
type Array struct {
	data []int
	size int
}

// NewArray create new array
func NewArray(capacity int) Array {

	return Array{data: make([]int, capacity)}

}

// NewDefaultArray create array with default size
func NewDefaultArray() Array {
	return Array{data: make([]int, 10)}

}

// GetArraySize get array size
func (a *Array) GetArraySize() int {

	return a.size
}

//GetCapacity  get the capacity of the array
func (a *Array) GetCapacity() int {
	return cap(a.data)
}

// IsEmpty check if the size is empty
func (a *Array) IsEmpty() bool {
	return a.size == 0
}

//PrintAarray will print all elements of the array
func (a *Array) PrintAarray() {
	fmt.Printf("now the size of array is: %d, items of the array is %d\n", a.size, a.data[:a.size])
}

func main() {

	arry1 := NewDefaultArray()
	arry1.PrintAarray()

	arry2 := NewArray(5)
	arry2.PrintAarray()
}
