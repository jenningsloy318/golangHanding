package array

import (
	"fmt"
	"testing"
)

//Array declare new Array
type Array struct {
	data []interface{}
	size int
}

// NewArray create new array
func NewArray(capacity int) Array {

	return Array{data: make([]interface{}, capacity)}

}

// NewDefaultArray create array with default size
func NewDefaultArray() Array {
	return Array{data: make([]interface{}, 10)}

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

//ToString will print the array with string
func (a *Array) ToString() string {
	var stringSlice []interface{}
	for i := 0; i < a.size; i++ {
		stringSlice = append(stringSlice, a.data[i])
	}

	return fmt.Sprintf("now the size of array is: %d, items of the array is %v\n", a.size, stringSlice)

}
func TestArray(t *testing.T) {

	arry1 := NewDefaultArray()
	t.Log(arry1.ToString())

	arry2 := NewArray(5)
	t.Log(arry2.ToString())
}
