package array

import (
	"fmt"
	"strings"
	"testing"
)

//Array declare new Array, and the array will implement the Stack interface
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
func (a *Array) GetArraySize() interface{} {

	return a.size
}

//GetCapacity  get the capacity of the array
func (a Array) GetCapacity() interface{} {
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
func (a *Array) Get(index int) interface{} {
	if index < 0 || index > a.size {
		fmt.Errorf("invalid index")
	}
	return a.data[index]
}

//GetLast the last element
func (a *Array) GetLast() interface{} {
	return a.Get(a.size - 1)
}

//GetFirst the first element
func (a *Array) GetFirst() interface{} {
	return a.Get(0)
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

	return fmt.Sprintf("stack size = %d,capacity =%d,stack Bottom: %v Top\n", a.size, a.GetCapacity(), stringSlice)
}

// Contains will check if the array contains certain element
func (a *Array) Contains(element interface{}) bool {

	for i := 0; i < a.size; i++ {
		if a.data[i] == element {
			return true
		}
	}
	return false
}

//Find check if the array contains certain element, return the index if exist, return -1 if not exist
func (a *Array) Find(element interface{}) (index int, err error) {

	for i := 0; i < a.size; i++ {
		if a.data[i] == element {
			return i, nil
		}
	}
	return -1, fmt.Errorf("Don't find the element")

}

//Remove will delete the element by index, and return the value of the element
func (a *Array) Remove(index int) interface{} {
	if index < 0 || index > a.size {
		fmt.Println("invalid index")
	}

	returnElement := a.data[index]
	for i := index + 1; i < a.size; i++ {
		a.data[i-1] = a.data[i]
	}

	a.size--
	//shrink the arary if the size is half of the capacity
	if a.size == cap(a.data)/4 && cap(a.data)/2 != 0 {
		a.resize(cap(a.data) / 2)
	}
	return returnElement
}

//RemoveLast remove the last element
func (a *Array) RemoveLast() interface{} {
	index := a.size - 1
	return a.Remove(index)

}

//RemoveFirst remove the first element
func (a *Array) RemoveFirst() interface{} {

	return a.Remove(0)
}

//RemoveElement remove the element
func (a *Array) RemoveElement(element interface{}) {
	index, err := a.Find(element)
	if err != nil {
		fmt.Errorf("don't find the element")
		return
	}

	a.Remove(index)
}

//RemoveElement remove the element
func (a *Array) resize(newCapacity int) {

	newData := make([]interface{}, newCapacity)
	for i := 0; i < a.size; i++ {
		newData[i] = a.data[i]
	}
	a.data = newData
}

//Stack define an interface
type Stack interface {
	GetSize() interface{}     // get size of the stack
	IsEmpty() bool            // check if the stack is empty
	Push(element interface{}) // push element to stack
	Pop() interface{}         // fetch the top element of the stack
	Peek() interface{}        // verify the top value
}

//GetSize implement interface Stack
func (a *Array) GetSize() interface{} {

	return a.GetArraySize()
}

//Push implement interface Stack
func (a *Array) Push(element interface{}) {

	a.AddLast(element)
}

//Pop implement interface Stack
func (a *Array) Pop() interface{} {

	return a.RemoveLast()
}

//Peek implement interface Stack
func (a *Array) Peek() interface{} {
	return a.GetLast()
}

func isValid(s string) bool {
	charSlice := strings.SplitAfter(s, "")
	var newAS = NewDefaultArray()
	for _, c := range charSlice {
		if c == "{" || c == "[" || c == "(" {
			newAS.Push(c)
			fmt.Println(newAS.ToString())
		} else {
			if newAS.IsEmpty() {
				return false
			}
			topChar := newAS.Pop()
			if c == "(" && topChar != ")" {
				return false
			}
			if c == "[" && topChar != "]" {
				return false
			}
			if c == "{" && topChar != "}" {
				return false
			}
		}

	}

	return newAS.IsEmpty()
}

func TestArayStack(t *testing.T) {

	char := "({})"
	t.Log(isValid(char))
}
