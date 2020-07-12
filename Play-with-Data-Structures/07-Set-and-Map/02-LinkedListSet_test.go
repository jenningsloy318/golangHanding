package linkedlistset

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

type linkedListSet interface {
	Add(element string)
	Contains(element string) bool
	remove(element string)
	GetSize() int
	IsEmpty() bool
}

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

func (l *LinkedList) AddWiwthIndex(index int, element interface{}) error {
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
	l.AddWiwthIndex(0, element)

}

func (l *LinkedList) Add(element interface{}) {
	if !l.Contains(element) {
		l.AddFirst(element)
	}

}

func (l *LinkedList) Contains(element interface{}) bool {

	for currentNode := l.dummyHead.next; currentNode != nil; currentNode = currentNode.next {
		if currentNode.element == element {
			return true
		}
	}
	return false
}

func (l *LinkedList) RemoveWithIndex(index int) (element interface{}) {
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

	return l.RemoveWithIndex(0)
}

func (l *LinkedList) RemoveLast() (element interface{}) {

	return l.RemoveWithIndex(l.size - 1)
}

func (l *LinkedList) Remove(element interface{}) {

	for previousNode := l.dummyHead; previousNode.next != nil; previousNode = previousNode.next {
		if previousNode.next.element == element {
			previousNode.next = previousNode.next.next
			l.size--
		}
	}

}

func (l *LinkedList) GetSize() int {
	return l.size
}

func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

func TestLinkedListSet(t *testing.T) {

	f1, _ := os.Open("pride-and-prejudice.txt")
	defer func() {
		if err := f1.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	var testSlice []string

	s := bufio.NewScanner(f1)
	for s.Scan() {
		s1 := strings.Fields(s.Text())
		testSlice = append(testSlice, s1...)
	}

	t.Logf("Total Words: %d", len(testSlice))

	newBst := NewLinkedList()
	for _, word := range testSlice {
		newBst.Add(word)
	}
	t.Logf("Total uniq Words: %d", newBst.GetSize())

}