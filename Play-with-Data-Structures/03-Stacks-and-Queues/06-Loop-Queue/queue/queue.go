package queue

import (
	"fmt"
	"strconv"
	"strings"
)

//Queue define a struct of Queue
type Queue struct {
	data  []int
	front int
	tail  int
	size  int
}

//NewQueue create new Queue
func NewQueue(capacity int) Queue {

	return Queue{
		data:  make([]int, capacity+1),
		front: 0,
		tail:  0,
		size:  0,
	}
}

//NewDefaultQueue will create new queue with 10
func NewDefaultQueue() Queue {
	return NewQueue(10)
}

//IsEmpty check if the queue is empty
func (q *Queue) IsEmpty() bool {

	return q.size == 0
}

//GetCapacity return the capacity of the queue
func (q *Queue) GetCapacity() int {
	//as this is loop queue, so underlying slice must have one element as nil, so the
	return cap(q.data) - 1
}

//GetSize return the size of the queue
func (q *Queue) GetSize() int {
	return q.size
}

//Enqueue will add an element to this queue
func (q *Queue) Enqueue(element int) {
	// check if the queue is full, as this is loop queue, so underlying slice must have one element as nil,
	// in normal queue when tail +1 as new tail when adding a new element, if this value is equal to front,  we treat queue is full , so there is no space for new element,
	// for loop queue , as maybe the front index is larger that tails, so use (tail + 1) % sliceCapacity  to get its  remainder as a new tail， if the remainder is equal to front, we treat it full
	sliceCapacity := cap(q.data)
	if (q.tail+1)%sliceCapacity == q.front {
		q.resize(q.GetCapacity() * 2)
		sliceCapacity = cap(q.data)
	}

	q.data[q.tail] = element

	//get the new tail as a remainder to its capacity , as a loop queue
	q.tail = (q.tail + 1) % sliceCapacity
	q.size++
}

//Dequeue get the head of the queue
func (q *Queue) Dequeue() int {
	sliceCapacity := cap(q.data)

	if q.IsEmpty() {
		fmt.Errorf("The queue is empty")
	}
	curentHead := q.data[q.front]
	// clear the currnt head, which is set the slice element to zero
	q.data[q.front] = 0
	q.front = (q.front + 1) % sliceCapacity
	q.size--

	if q.size == q.GetCapacity()/4 && q.GetCapacity()/2 != 0 {
		q.resize(q.GetCapacity() / 2)
	}
	return curentHead
}

//ToString will print the queue as  string
func (q *Queue) ToString() string {

	//we can loop the queue from front to tail , and convert the front and tail to actual slice index
	sliceCapacity := cap(q.data)
	var queueSlice []string
	for i := q.front; i != q.tail; i = (i + 1) % sliceCapacity {
		queueSlice = append(queueSlice, strconv.Itoa(q.data[i]))

	}
	return fmt.Sprintf("queue size = %d,capacity = %d,Queue Head: [%s] Tail\n", q.size, q.GetCapacity(), strings.Join(queueSlice[:], ","))
}

func (q *Queue) resize(capacity int) {
	oldSliceCapacity := cap(q.data)
	newData := make([]int, capacity+1)
	for i := 0; i < q.size; i++ {
		// i=0: data[front] ==> newData[0]
		// i=1: data[front +1]  ==>newData[1]
		// as a loop queue, must get remainder as next front
		newData[i] = q.data[(q.front+i)%oldSliceCapacity]
	}
	q.data = newData
	q.front = 0
	q.tail = q.size

}
