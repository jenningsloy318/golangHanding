package array

import (
	"fmt"
	"strings"
)

//Array declare new Array, and the array will implement the Stack interface
type Array struct {
	data []rune
	size int
}

// NewArray create new array
func NewArray(capacity int) Array {

	return Array{data: make([]rune, capacity), size: 0}

}

// NewDefaultArray create array with default size
func NewDefaultArray() Array {
	return Array{data: make([]rune, 10), size: 0}

}

// GetArraySize get array size
func (a *Array) GetArraySize() int {

	return a.size
}

//GetCapacity  get the capacity of the array
func (a Array) GetCapacity() int {
	return cap(a.data)
}

// IsEmpty check if the size is empty, and also implemt the Stack interface
func (a *Array) IsEmpty() bool {
	return a.size == 0
}

//PrintAarray will print all elements of the array
func (a *Array) PrintAarray() {
	fmt.Printf("now the size of array is: %d, capacity is %d, items of the array is %d\n", a.size, a.GetCapacity(), a.data[:a.size])
}

// AddLast will add an item to the last
func (a *Array) AddLast(element rune) {

	//if a.size == cap(a.data) {
	//	fmt.Errorf("Array is full, AddLast can't add element to the Array")
	//}
	//a.data[a.size] = element

	a.Add(a.size, element)
}

// AddFirst will add an item to the first
func (a *Array) AddFirst(element rune) {
	a.Add(0, element)
}

// Add will add element at the index postion
func (a *Array) Add(index int, element rune) {

	if index < 0 || index > a.size {
		fmt.Errorf("invalid index")
	}

	//increase the arary if the array is full

	if a.size == cap(a.data) {
		a.resize(2 * a.size)
	}
	// move t.data[i] ---> t.data[i+1]
	for i := a.size - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}

	a.data[index] = element
	a.size++

}

//Get will the value for index
func (a *Array) Get(index int) rune {
	if index < 0 || index > a.size {
		fmt.Errorf("invalid index")
	}
	return a.data[index]
}

//GetLast the last element
func (a *Array) GetLast() rune {
	return a.Get(a.size - 1)
}

//GetFirst the first element
func (a *Array) GetFirst() rune {
	return a.Get(0)
}

//Set will set value for index
func (a *Array) Set(index int, element rune) {
	if index < 0 || index > a.size {
		fmt.Println("invalid index")
	}
	a.data[index] = element
}

//ToString will print the array with string
func (a *Array) ToString() string {
	var stringSlice []string
	for i := 0; i < a.size; i++ {
		stringSlice = append(stringSlice, string(a.data[i]))
	}

	return strings.Join(stringSlice[:], ",")

}

// Contains will check if the array contains certain element
func (a *Array) Contains(element rune) bool {

	for i := 0; i < a.size; i++ {
		if a.data[i] == element {
			return true
		}
	}
	return false
}

//Find check if the array contains certain element, return the index if exist, return -1 if not exist
func (a *Array) Find(element rune) (index int, err error) {

	for i := 0; i < a.size; i++ {
		if a.data[i] == element {
			return i, nil
		}
	}
	return -1, fmt.Errorf("Don't find the element")

}

//Remove will delete the element by index, and return the value of the element
func (a *Array) Remove(index int) rune {
	if index < 0 || index > a.size {
		fmt.Println("invalid index")
	}

	returnElement := a.data[index]
	for i := index + 1; i < a.size; i++ {
		a.data[i-1] = a.data[i]
	}

	a.size--
	//shrink the arary if the size is half of the capacity
	if a.size == cap(a.data)/2 {
		a.resize(a.size)
	}
	return returnElement
}

//RemoveLast remove the last element
func (a *Array) RemoveLast() rune {
	index := a.size - 1
	return a.Remove(index)

}

//RemoveFirst remove the first element
func (a *Array) RemoveFirst() rune {

	return a.Remove(0)
}

//RemoveElement remove the element
func (a *Array) RemoveElement(element rune) {
	index, err := a.Find(element)
	if err != nil {
		fmt.Errorf("don't find the element")
		return
	}

	a.Remove(index)
}

//RemoveElement remove the element
func (a *Array) resize(newCapacity int) {

	newData := make([]rune, newCapacity)
	for i := 0; i < a.size; i++ {
		newData[i] = a.data[i]
	}
	a.data = newData
}

//GetSize implement interface Stack
func (a *Array) GetSize() int {

	return a.GetArraySize()
}

//Push implement interface Stack
func (a *Array) Push(element rune) {

	a.AddLast(element)
}

//Pop implement interface Stack
func (a *Array) Pop() rune {

	return a.RemoveLast()
}

//Peek implement interface Stack
func (a *Array) Peek() rune {
	return a.GetLast()
}
