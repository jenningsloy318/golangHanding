package set

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

type bstSet interface {
	add(element string)
	contains(element string) bool
	remove(element string)
	getSize() int
	isEmpty() bool
}

type Node struct {
	value string
	left  *Node
	right *Node
}

func NewNode(value string) *Node {
	return &Node{
		value: value,
	}
}

type BsTree struct {
	root *Node
	size int
}

func NewBsTree() *BsTree {
	return &BsTree{
		root: nil,
		size: 0,
	}
}

func (b *BsTree) Size() int {
	return b.size
}

func (b *BsTree) IsEmpty() bool {
	return b.size == 0
}

func (b *BsTree) Add(value string) {
	b.root = b.add(b.root, value)

}

func (b *BsTree) add(node *Node, value string) *Node {
	if node == nil {
		b.size++
		return NewNode(value)
	}
	if value < node.value {
		node.left = b.add(node.left, value)
	}

	if value > node.value {

		node.right = b.add(node.right, value)
	}

	return node

}

func (b *BsTree) Contains(value string) bool {
	return contains(b.root, value)
}

func contains(node *Node, value string) bool {

	if node == nil {
		return false
	}

	if node.value == value {
		return true
	} else if value < node.value {
		return contains(node.left, value)
	} else // value > node.value
	{
		return contains(node.right, value)
	}

}

func preOrderTraverse(node *Node) {

	if node == nil {
		return
	}

	fmt.Println(node.value)
	preOrderTraverse(node.left)
	preOrderTraverse(node.right)

}

func (b *BsTree) InOrderTraverse() {
	inOrderTraverse(b.root)

}
func inOrderTraverse(node *Node) {
	if node == nil {
		return
	}

	inOrderTraverse(node.left)
	fmt.Println(node.value)
	inOrderTraverse(node.right)

}

func (b *BsTree) PostOrderTraverse() {
	postOrderTraverse(b.root)

}

func postOrderTraverse(node *Node) {
	if node == nil {
		return
	}
	postOrderTraverse(node.left)
	postOrderTraverse(node.right)
	fmt.Println(node.value)

}

func TestBstreeBasic(t *testing.T) {

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

	newBst := NewBsTree()
	for _, word := range testSlice {
		newBst.Add(word)
	}
	t.Logf("Total uniq Words: %d", newBst.Size())

}
