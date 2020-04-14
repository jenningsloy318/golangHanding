package queue

//Queue define an interface
type Queue interface {
	GetSize() int        // get size of the stack
	IsEmpty() bool       // check if the stack is empty
	Enqueue(element int) // push element to stack
	Dequeue() int        // fetch the top element of the stack
	GetFront() int       // verify the top value
}
