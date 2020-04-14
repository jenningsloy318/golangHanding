package stack

//Stack define an interface
type Stack interface {
	GetSize() int      // get size of the stack
	IsEmpty() bool     // check if the stack is empty
	Push(element rune) // push element to stack
	Pop() rune         // fetch the top element of the stack
	Peek() rune        // verify the top value
}
