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

	return Array{data: make([]interface{}, capacity), size: 0}

}

// NewDefaultArray create array with default size
func NewDefaultArray() Array {
	return Array{data: make([]interface{}, 10), size: 0}

}

// GetArraySize get array size
func (a *Array) GetArraySize() int {

	return a.size
}

//GetCapacity  get the capacity of the array
func (a Array) GetCapacity() int {
	return cap(a.data)
}

// IsEmpty check if the size is empty
func (a *Array) IsEmpty() bool {
	return a.size == 0
}

// AddLast will add an item to the last
func (a *Array) AddLast(element interface{}) {

	//if a.size == cap(a.data) {
	//	fmt.Errorf("Array is full, AddLast can't add element to the Array")
	//}
	//a.data[a.size] = element

	a.Add(a.size, element)
}

// AddFirst will add an item to the first
func (a *Array) AddFirst(element interface{}) {
	a.Add(0, element)
}

// Add will add element at the index postion
func (a *Array) Add(index int, element interface{}) {

	if a.size == cap(a.data) {
		fmt.Errorf("Array is full, AddLast can't add element to the Array")
	}

	if index < 0 || index > a.size {
		fmt.Errorf("invalid index")
	}

	// move t.data[i] ---> t.data[i+1]
	for i := a.size - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}

	a.data[index] = element
	a.size++

}

//Get will the value for index
func (a *Array) Get(index int) interface{} {
	if index < 0 || index > a.size {
		fmt.Errorf("invalid index")
	}
	return a.data[index]
}

//Set will set value for index
func (a *Array) Set(index int, element interface{}) {
	if index < 0 || index > a.size {
		fmt.Println("invalid index")
	}
	a.data[index] = element
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

	array := NewDefaultArray()
	t.Log(array.ToString())

	array.AddLast(6)
	t.Log(array.ToString())

	array.Add(0, 7)
	t.Log(array.ToString())

	array.Add(1, 8)
	t.Log(array.ToString())

	array.Set(2, 1024)
	t.Logf("String list of the array elements is: %s \n", array.ToString())
	t.Logf("The value of %d element in array is: %d\n", 2, array.Get(2))
}
