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

func (l *LinkedList) Add(element interface{}, index int) error {
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
	l.Add(element, 0)

}
func (l *LinkedList) AddLast(element interface{}) {
	l.Add(element, l.size)
}

func (l *LinkedList) ToString() string {
	var allStrings []string
	for node := l.dummyHead; node != nil; node = node.next {
		allStrings = append(allStrings, fmt.Sprintf("%v", node.element))
	}
	return strings.Join(allStrings[:], ",")
}
func TestLinkedList(t *testing.T) {
	newList := NewLinkedList()
	newList.AddFirst(1)
	t.Log(newList.ToString())
	newList.AddFirst(2)
	t.Log(newList.ToString())
	newList.Add(3, 1)
	t.Log(newList.ToString())

}
