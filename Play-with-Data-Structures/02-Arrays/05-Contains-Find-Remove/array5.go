package main

import (
	"fmt"
	"strconv"
	"strings"
)

//Array declare new Array
type Array struct {
	data []int
	size int
}

// NewArray create new array
func NewArray(capacity int) Array {

	return Array{data: make([]int, capacity), size: 0}

}

// NewDefaultArray create array with default size
func NewDefaultArray() Array {
	return Array{data: make([]int, 10), size: 0}

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

//PrintAarray will print all elements of the array
func (a *Array) PrintAarray() {
	fmt.Printf("now the size of array is: %d, items of the array is %d\n", a.size, a.data[:a.size])
}

// AddLast will add an item to the last
func (a *Array) AddLast(element int) {

	//if a.size == cap(a.data) {
	//	fmt.Errorf("Array is full, AddLast can't add element to the Array")
	//}
	//a.data[a.size] = element

	a.Add(a.size, element)
}

// AddFirst will add an item to the first
func (a *Array) AddFirst(element int) {
	a.Add(0, element)
}

// Add will add element at the index postion
func (a *Array) Add(index int, element int) {

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
func (a *Array) Get(index int) int {
	if index < 0 || index > a.size {
		fmt.Errorf("invalid index")
	}
	return a.data[index]
}

//Set will set value for index
func (a *Array) Set(index int, element int) {
	if index < 0 || index > a.size {
		fmt.Println("invalid index")
	}
	a.data[index] = element
}

//ToString will print the array with string
func (a *Array) ToString() string {
	var stringSlice []string
	for i := 0; i < a.size; i++ {
		stringSlice = append(stringSlice, strconv.Itoa(a.data[i]))
	}

	return strings.Join(stringSlice[:], ",")

}

// Contains will check if the array contains certain element
func (a *Array) Contains(element int) bool {

	for i := 0; i < a.size; i++ {
		if a.data[i] == element {
			return true
		}
	}
	return false
}

//Find check if the array contains certain element, return the index if exist, return -1 if not exist
func (a *Array) Find(element int) (index int, err error) {

	for i := 0; i < a.size; i++ {
		if a.data[i] == element {
			return i, nil
		}
	}
	return -1, fmt.Errorf("Don't find the element")

}

//Remove will delete the element by index, and return the value of the element
func (a *Array) Remove(index int) int {
	if index < 0 || index > a.size {
		fmt.Println("invalid index")
	}

	returnElement := a.data[index]
	for i := index + 1; i < a.size; i++ {
		a.data[i-1] = a.data[i]
	}

	a.size--
	return returnElement
}

//RemoveLast remove the last element
func (a *Array) RemoveLast() int {
	index := a.size - 1
	return a.Remove(index)

}

//RemoveFirst remove the first element
func (a *Array) RemoveFirst() int {

	return a.Remove(0)
}

//RemoveElement remove the element
func (a *Array) RemoveElement(element int) {
	index, err := a.Find(element)
	if err != nil {
		fmt.Errorf("don't find the element")
		return
	}

	a.Remove(index)
}

func main() {

	array := NewDefaultArray()
	array.PrintAarray()

	array.AddLast(6)
	array.PrintAarray()

	array.Add(0, 7)
	array.PrintAarray()

	array.Add(1, 8)
	array.PrintAarray()

	array.Set(2, 1024)
	fmt.Printf("String list of the array elements is: %s \n", array.ToString())
	fmt.Printf("The value of %d element in array is: %d\n", 2, array.Get(2))
	array.Add(3, 2048)

	array.PrintAarray()

	array.RemoveLast()
	array.Add(3, 3096)
	array.PrintAarray()

	array.RemoveElement(3096)

	array.PrintAarray()

}
