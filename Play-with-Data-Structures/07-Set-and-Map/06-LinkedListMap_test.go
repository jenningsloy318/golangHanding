package linkedlistmap

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

type nMap interface {
	Add(key string, value int)
	Remove(key string) int
	Contains(key string) bool
	Get(key string) int
	Set(key, string, value int)
	GetSize() int
	IsEmpty() bool
}

type Node struct {
	Key   string
	Value int
	Next  *Node
}

func NewNode(key string, value int, next *Node) *Node {
	return &Node{
		Key:   key,
		Value: value,
		Next:  next,
	}
}

func NewEmptyNode() *Node {
	return &Node{}
}

func NewKeyNode(key string) *Node {
	return &Node{
		Key: key,
	}
}

func (n *Node) ToString() string {
	return fmt.Sprintf("%d", n.Value)
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

func (l *LinkedList) GetSize() int {
	return l.size
}

func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList) getNode(key string) *Node {

	for curNode := l.dummyHead.Next; curNode != nil; curNode = curNode.Next {
		if curNode.Key == key {
			return curNode
		}

	}
	return nil
}

func (l *LinkedList) Contains(key string) bool {
	return l.getNode(key) != nil

}

func (l *LinkedList) Get(key string) int {

	curNode := l.getNode(key)
	if curNode != nil {
		return curNode.Value
	}
	return -1

}
func (l *LinkedList) Set(key string, value int) {

	curNode := l.getNode(key)
	if curNode != nil {
		curNode.Value = value
		return
	}
	fmt.Errorf("Key %s doesn't exist", key)
}

func (l *LinkedList) Add(key string, value int) {

	node := l.getNode(key)
	if node == nil {
		l.dummyHead.Next = NewNode(key, value, l.dummyHead.Next)
		l.size++
		return
	} else {
		node.Value = value
	}

}

func (l *LinkedList) Remove(key string) int {

	for preNode := l.dummyHead; preNode.Next != nil; preNode = preNode.Next {
		if preNode.Next.Key == key {
			returnNode := preNode.Next
			preNode.Next = preNode.Next.Next
			l.size--
			return returnNode.Value
		}

	}
	return -1
}

func TestLinkedList(t *testing.T) {
	file, _ := os.Open("pride-and-prejudice.txt")
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	var testSlice []string

	s := bufio.NewScanner(file)
	for s.Scan() {
		s1 := strings.Fields(s.Text())
		testSlice = append(testSlice, s1...)
	}

	t.Logf("Total Words: %d", len(testSlice))

	newList := NewLinkedList()
	t1 := time.Now()
	for _, word := range testSlice {

		if newList.Contains(word) {
			newList.Set(word, newList.Get(word)+1)
		} else {
			newList.Add(word, 1)
		}

	}
	duration := time.Since(t1)
	t.Logf("Total uniq Words: %d, process time %fs ", newList.GetSize(), duration.Seconds())
	t.Logf("Frequence of  to is : %d.\n", newList.Get("to"))
	t.Log("Remove word to .")
	newList.Remove("to")
	t.Logf("Now the frequence of  to is : %d.\n", newList.Get("to"))

}
