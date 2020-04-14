package queue

//Queue define a struct of Queue
type Queue struct {
	data  []int
	front int
	tail  int
	size  int
}

//NewQueue create new Queue
func NewQueue(int capacity) Queue {

	return Queue{
		data:  make([]int, capacity+1),
		front: 0,
		tail:  0,
		size:  0,
	}
}

//NewDefaultQueue will create new queue with 10
func NewDefaultQueue() Queue {
	NewQueue(10)
}

//IsEmpty check if the queue is empty
func (q *Queue) IsEmpty() bool {

	return q.size == 0
}

//GetCapacity return the capacity of the queue
func (q *Queue) GetCapacity() int {
	return cap(q.data) - 1
}

//GetSize return the size of the queue
func (q *Queue) GetSize() int {
	return q.size
}
