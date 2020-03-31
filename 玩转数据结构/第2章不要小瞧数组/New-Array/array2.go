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
func (t Array) GetArraySize() int {

	return t.size
}

//GetCapacity  get the capacity of the array
func (t Array) GetCapacity() int {
	return cap(t.data)
}

// IsEmpty check if the size is empty
func (t Array) IsEmpty() bool {
	return t.size == 0
}
func main() {

	arry1 := NewDefaultArray()
	fmt.Println(arry1.GetArraySize())
	fmt.Println(arry1.GetCapacity())

	arry2 := NewArray(5)
	fmt.Println(arry2.GetArraySize())
	fmt.Println(arry2.GetCapacity())
}
