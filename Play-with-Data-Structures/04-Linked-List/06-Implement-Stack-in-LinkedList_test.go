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
	dummyHead *Node
	size      int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		dummyHead: NewEmptyNode(),
		size:      0,
	}
}

func (l *LinkedList) Add(index int, element interface{}) error {
	if index < 0 || index > l.size {
		return fmt.Errorf("Invalid index")
	}
	previousNode := l.dummyHead
	for i := 0; i < index; i++ {
		previousNode = previousNode.next
	}
	//newNode := NewSingleNode(element)
	//newNode.next = previousNode.next
	//previousNode.next = newNode

	previousNode.next = NewNode(element, previousNode.next)
	l.size++
	return nil
}

func (l *LinkedList) AddFirst(element interface{}) {
	l.Add(0, element)

}
func (l *LinkedList) AddLast(element interface{}) {
	l.Add(l.size-1, element)
}

func (l *LinkedList) ToString() string {
	var allStrings []string
	for currentNode := l.dummyHead.next; currentNode != nil; currentNode = currentNode.next {
		allStrings = append(allStrings, fmt.Sprintf("%v", currentNode.element))
	}
	return strings.Join(allStrings[:], ",")
}

func (l *LinkedList) Get(index int) (element interface{}) {

	if index < 0 || index > l.size {
		return fmt.Errorf("Invalid index")
	}

	currentNode := l.dummyHead.next
	for i := 0; i < index; i++ {
		currentNode = currentNode.next
	}

	return currentNode.element
}

func (l *LinkedList) GetFirst() (element interface{}) {

	return l.Get(0)
}

func (l *LinkedList) GetLast() (element interface{}) {

	return l.Get(l.size - 1)
}

func (l *LinkedList) Set(index int, element interface{}) {

	if index < 0 || index > l.size {
		fmt.Errorf("Invalid index")
	}

	currentNode := l.dummyHead.next
	for i := 0; i < index; i++ {
		currentNode = currentNode.next
	}

	currentNode.element = element
}

func (l *LinkedList) Contains(element interface{}) bool {

	for currentNode := l.dummyHead.next; currentNode != nil; currentNode = currentNode.next {
		if currentNode.element == element {
			return true
		}
	}
	return false
}
func (l *LinkedList) Remove(index int) (element interface{}) {
	if index < 0 || index > l.size {
		fmt.Errorf("Invalid index")
	}
	previousNode := l.dummyHead
	for i := 0; i < index; i++ {
		previousNode = previousNode.next
	}

	curNode := previousNode.next
	previousNode.next = curNode.next
	curNode.next = nil

	l.size--
	return curNode.element

}

func (l *LinkedList) RemoveFirst() (element interface{}) {

	return l.Remove(0)
}

func (l *LinkedList) RemoveLast() (element interface{}) {

	return l.Remove(l.size - 1)
}

//Stack define an interface
type Stack interface {
	GetSize() int                // get size of the stack
	IsEmpty() bool               // check if the stack is empty
	Push(element interface{})    // push element to stack
	Pop() (element interface{})  // fetch the top element of the stack
	Peek() (element interface{}) // verify the top value
}

func (l *LinkedList) GetSize() int {

	return l.size
}

func (l *LinkedList) IsEmpty() bool {

	return l.size == 0
}

func (l *LinkedList) Push(element interface{}) {

	l.AddFirst(element)
}

func (l *LinkedList) Pop() {

	l.RemoveFirst()
}

func (l *LinkedList) Peek() (element interface{}) {

	return l.GetFirst()
}

func TestLinkedList(t *testing.T) {
	newList := NewLinkedList()

	for i := 0; i < 5; i++ {
		newList.AddFirst(i)
		t.Log(newList.ToString())
	}

	newList.Add(2, 4444)
	t.Log(newList.ToString())

	t.Logf("list contains 3: %t", newList.Contains(3))

	t.Logf("list contains 4444: %t", newList.Contains(4444))

	newList.Remove(2)
	t.Log(newList.ToString())

	newList.RemoveFirst()
	t.Log(newList.ToString())

	newList.RemoveLast()
	t.Log(newList.ToString())

	//
	newList.Push("aaa")
	t.Logf("Push element %s to stack, now the stack is %v", "aaa", newList.ToString())

	t.Log(newList.Peek())

	newList.Pop()
	t.Logf("Pop element from stack, now the stack is %v", newList.ToString())

}
