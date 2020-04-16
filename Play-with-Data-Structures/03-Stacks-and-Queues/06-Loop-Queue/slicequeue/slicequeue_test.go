package slicequeue

import (
	"fmt"
	"testing"
)

//SliceQueue define a stack
type SliceQueue []interface{}

//NewSliceQueue create new SliceQueue
func NewSliceQueue() SliceQueue {
	return make(SliceQueue, 0)
}

//Push  an element to SliceQueue
func (sq *SliceQueue) Enqueue(element interface{}) {

	*sq = append(*sq, element)
	fmt.Printf("SliceQueue: add element: %v, now the SliceQueue is %v \n", element, *sq)
}

//Dequeue will get the pop the front element from SliceQueue
func (sq *SliceQueue) Dequeue() (element interface{}) {
	length := len(*sq)
	if length == 0 {
		fmt.Printf("No element in the queue")
		return nil
	}

	ret := (*sq)[0]

	*sq = (*sq)[1:]

	fmt.Printf("SliceQueue: Pop element: %v, now the SliceQueue is %#v \n", ret, *sq)

	return ret
}

//GetFront get the front element
func (sq *SliceQueue) GetFront() (element interface{}) {
	return (*sq)[0]
}

//ToString print the string  list of the SliceQueue
func (sq *SliceQueue) ToString() string {

	return fmt.Sprintf("Front %#v Tail", *sq)
}

//IsEmpty
func (sq *SliceQueue) IsEmpty() bool {
	return len(*sq) == 0
}

//TestStack test
func TestStack(t *testing.T) {

	var loopQueue = NewSliceQueue()

	for i := 0; i <= 15; i++ {
		loopQueue.Enqueue(i)
		t.Log(loopQueue.ToString())
		if i%3 == 2 {
			loopQueue.Dequeue()
			t.Log(loopQueue.ToString())
		}
	}
}
