package arraystack

import (
	"fmt"
	"strings"
	"testing"
)

//ArrayStack define a stack
type ArrayStack []interface{}

//NewArrayStack create new ArrayStack
func NewArrayStack() ArrayStack {
	return make(ArrayStack, 0)
}

//Push  an element to ArrayStack
func (as *ArrayStack) Push(element interface{}) {

	*as = append(*as, element)
	fmt.Printf("ArrayStack: Push element: %v, now the arraystack is %v \n", element, *as)
}

//Pop will get the pop the top element from ArrayStack
func (as *ArrayStack) Pop() (element interface{}) {
	length := len(*as)
	if length == 0 {
		fmt.Errorf("The stack is empty")
		return nil
	}

	ret := (*as)[length-1]
	if length == 1 {
		*as = make(ArrayStack, 0)
	} else {
		*as = (*as)[:length-1]
	}
	fmt.Printf("ArrayStack: Pop element: %v, now the arraystack is %#v \n", ret, *as)

	return ret
}

//Peek get the top element
func (as *ArrayStack) Peek() (element interface{}) {
	return (*as)[len(*as)-1]
}

//ToString print the string  list of the ArrayStack
func (as *ArrayStack) ToString() string {

	return fmt.Sprintf("Bottom %#v Top", *as)
}

//IsEmpty
func (as *ArrayStack) IsEmpty() bool {
	return len(*as) == 0
}

//TestStack test
func TestStack(t *testing.T) {

	as := NewArrayStack()
	s := "[{()}]"
	for _, c := range strings.SplitAfter(s, "") {
		if c == "{" || c == "[" || c == "(" {
			as.Push(c)
			t.Log(as.ToString(), c)
		} else {
			if as.IsEmpty() {
				t.Log("false")
			}

			topChar, _ := as.Pop().(string)
			t.Log(topChar, c)

			if c == "(" && topChar != ")" {
				t.Log("false")
			}

			if c == "[" && topChar != "]" {
				t.Log("false")
			}

			if c == "{" && topChar != "}" {
				t.Log("false")
			}

		}

	}

	if as.IsEmpty() {
		t.Log("true")
	}
}
