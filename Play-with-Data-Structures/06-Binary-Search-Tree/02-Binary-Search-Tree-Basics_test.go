package bst

import "testing"

type Node struct {
	value interface{}
	left  *Node
	right *Node
}

func NewNode(value interface{}) *Node {
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

func TestBstreeBasic(t *testing.T) {

	newNode := NewNode(5)
	t.Logf("new node is : %v", *newNode)

	newBsTree := NewBsTree()

	t.Logf("new bs tree is : %v", *newBsTree)

}
