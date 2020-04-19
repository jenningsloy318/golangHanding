package stack

//Stack define an interface
type Stack interface {
	GetSize() int        // get size of the stack
	IsEmpty() bool       // check if the stack is empty
	Push(element string) // push element to stack
	Pop() string         // fetch the top element of the stack
	Peek() string        // verify the top value
}
