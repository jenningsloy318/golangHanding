package slicestack

import (
	"fmt"
	"testing"
)

//SliceStack define a stack
type SliceStack []interface{}

//NewSliceStack create new SliceStack
func NewSliceStack() SliceStack {
	return make(SliceStack, 0)
}

//Push  an element to SliceStack
func (ss *SliceStack) Push(element interface{}) {

	*ss = append(*ss, element)
	fmt.Printf("SliceStack: Push element: %v, now the arraystack is %v \n", element, *ss)
}

//Pop will get the pop the top element from SliceStack
func (ss *SliceStack) Pop() (element interface{}) {
	length := len(*ss)
	if length == 0 {
		return nil
	}

	ret := (*ss)[length-1]
	if length == 1 {
		*ss = make(SliceStack, 0)
	} else {
		*ss = (*ss)[:length-1]
	}
	fmt.Printf("SliceStack: Pop element: %v, now the arraystack is %#v \n", ret, *ss)

	return ret
}

//Peek get the top element
func (ss *SliceStack) Peek() (element interface{}) {
	return (*ss)[len(*ss)-1]
}

//ToString print the string  list of the SliceStack
func (ss *SliceStack) ToString() string {

	return fmt.Sprintf("Bottom %#v Top", *ss)
}

//IsEmpty
func (ss *SliceStack) IsEmpty() bool {
	return len(*ss) == 0
}

//TestStack test
func TestStack(t *testing.T) {

	ss := NewSliceStack()
	for i := 1; i < 15; i++ {
		ss.Push(i)
		t.Log("ss:", ss.ToString())

		if i%3 == 1 {
			ss.Pop()
			t.Log("ss:", ss.ToString())
		}

	}

	pp := []string{"a", "b", "c", "d"}
	for _, e := range pp {
		ss.Push(e)
		t.Log("ss:", ss.ToString())
	}
}
