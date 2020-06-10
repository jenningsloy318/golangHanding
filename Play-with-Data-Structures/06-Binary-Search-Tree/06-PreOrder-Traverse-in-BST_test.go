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
	b.root = b.add(b.root, value)

}

func (b *BsTree) add(node *Node, value int) *Node {
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

func (b *BsTree) Contains(value int) bool {
	return contains(b.root, value)
}

func contains(node *Node, value int) bool {

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

func (b *BsTree) PreOrderTraverse() {
	preOrderTraverse(b.root)

}

func preOrderTraverse(node *Node) {

	if node == nil {
		return
	}

	fmt.Println(node.value)
	preOrderTraverse(node.left)
	preOrderTraverse(node.right)

}
func TestBstreeBasic(t *testing.T) {

	newBsTree := NewBsTree()

	numbers := []int{33, 21, 77, 99, 36, 15, 76}
	for _, number := range numbers {
		newBsTree.Add(number)
	}
	newBsTree.PreOrderTraverse()
}
