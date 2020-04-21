package array

import (
	"fmt"
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
func (a *Array) GetArraySize() int {

	return a.size
}

//GetCapacity  get the capacity of the array
func (a Array) GetCapacity() int {
	return cap(a.data)
}

// IsEmpty check if the size is empty, and also implemt the Queue interface
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

	return fmt.Sprintf("queue size = %d,capacity =%d,Queue Head: %v Tail\n", a.size, a.GetCapacity(), stringSlice)

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

//Queue define an interface
type Queue interface {
	GetSize() int                // get size of the stack
	IsEmpty() bool               // check if the stack is empty
	Enqueue(element interface{}) // push element to stack
	Dequeue() int                // fetch the top element of the stack
	GetFront() int               // verify the top value
}

//GetSize implement interface Queue
func (a *Array) GetSize() int {

	return a.GetArraySize()
}

//Enqueue implement interface Queue
func (a *Array) Enqueue(element interface{}) {

	a.AddLast(element)
}

//Dequeue implement interface Queue
func (a *Array) Dequeue() interface{} {

	return a.RemoveFirst()
}

//GetFront implement interface Queue
func (a *Array) GetFront() interface{} {
	return a.GetFirst()
}

func TestArrayQueue(t *testing.T) {

	var newAS = NewDefaultArray()

	for i := 0; i <= 15; i++ {
		newAS.Enqueue(i)
		t.Log(newAS.ToString())
		if i%3 == 2 {
			newAS.Dequeue()
			t.Log(newAS.ToString())
		}
	}
}
