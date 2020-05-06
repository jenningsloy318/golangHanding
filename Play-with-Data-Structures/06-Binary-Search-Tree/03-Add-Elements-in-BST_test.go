package bst

import (
	"fmt"
	"testing"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func NewNode(value int) *Node {
	return &Node{
		value: value,
	}
}

func ToString(node *Node) string {
	if node == nil {
		return ""
	}
	if node.left == nil {
		return fmt.Sprintf("left:0 - value: %v - right: %v\n", node.value, ToString(node.right))
	}
	if node.right == nil {
		return fmt.Sprintf("left:%v - value: %v - right: 0\n", ToString(node.left), node.value)
	}
	return fmt.Sprintf("left:%v - value: %v - right: %v\n", ToString(node.left), node.value, ToString(node.right))

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

func (b *BsTree) Add(value int) {
	if b.root == nil {
		b.root = NewNode(value)
		b.size++
	}

	add(b.root, value)
	b.size++
}

func add(node *Node, value int) {
	if node.value == value {
		return
	}
	newAddNode := NewNode(value)
	if value < node.value {
		if node.left == nil {
			node.left = newAddNode
		} else {
			add(node.left, value)
		}
		return
	}
	if value > node.value {
		if node.right == nil {
			node.right = newAddNode
		} else {
			add(node.right, value)
		}
		return
	}
}
func TestBstreeBasic(t *testing.T) {

	newNode := NewNode(5)
	t.Logf("new node is : %v", *newNode)

	newBsTree := NewBsTree()

	t.Logf("new bs tree is : %v", *newBsTree)

	newBsTree.Add(55)
	t.Logf("new bs tree is : %v", ToString(newBsTree.root))
	newBsTree.Add(65)
	t.Logf("new bs tree is : %v", ToString(newBsTree.root))
	newBsTree.Add(44)
	t.Logf("new bs tree is : %v", ToString(newBsTree.root))

}
