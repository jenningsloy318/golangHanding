package linkedlist

import (
	"fmt"
	"strings"
	"testing"
)

type Node struct {
	element interface{}
	next    *Node
}

func NewNode(element interface{}, next *Node) *Node {
	return &Node{
		element: element,
		next:    next,
	}
}

func NewEmptyNode() *Node {
	return NewNode(nil, nil)
}

func NewSingleNode(element interface{}) *Node {
	return NewNode(element, nil)

}

func (n *Node) ToString() string {
	return fmt.Sprintf("%v", n.element)
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (l *LinkedList) ToString() string {

	var allStrings []string

	for currentNode := l.head; currentNode != nil; currentNode = currentNode.next {
		allStrings = append(allStrings, fmt.Sprintf(" %v ->", currentNode.element))
	}
	return strings.Join(allStrings[:], "")
}

//Queue define an interface
type Queue interface {
	GetSize() int                    // get size of the stack
	IsEmpty() bool                   // check if the stack is empty
	Enqueue(element interface{})     // push element to stack
	Dequeue() (element interface{})  // fetch the top element of the stack
	GetFront() (element interface{}) // verify the top value
}

func (l *LinkedList) GetSize() int {

	return l.size
}

func (l *LinkedList) IsEmpty() bool {

	return l.size == 0
}

func (l *LinkedList) Enqueue(element interface{}) {

	if l.IsEmpty() {
		l.tail = NewSingleNode(element)
		l.head = l.tail
	} else {
		l.tail.next = NewSingleNode(element)
		l.tail = l.tail.next
	}

	l.size++
}

func (l *LinkedList) Dequeue() (element interface{}) {
	if l.IsEmpty() {
		fmt.Errorf("Can't dequeue, the queue is empty")
	}

	retNode := l.head

	if l.GetSize() == 1 {
		l.head = nil
		l.tail = nil
	} else {
		l.head = l.head.next

	}
	retNode.next = nil
	l.size--
	return retNode.element

}

func TestLinkedList(t *testing.T) {
	newList := NewLinkedList()

	for i := 1; i < 15; i++ {
		newList.Enqueue(i)
		t.Logf("after enqueued: %v\n", newList.ToString())
		if i%3 == 2 {
			newList.Dequeue()
			t.Logf("after dequeued: %v\n", newList.ToString())
		}
	}

}
