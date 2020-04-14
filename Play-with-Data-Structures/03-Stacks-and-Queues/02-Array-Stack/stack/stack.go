package stack

//Stack define an interface
type Stack interface {
	GetSize() int     // get size of the stack
	IsEmpty() bool    // check if the stack is empty
	Push(element int) // push element to stack
	Pop() int         // fetch the top element of the stack
	Peek() int        // verify the top value
}
