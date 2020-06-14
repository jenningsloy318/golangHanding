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

// to use NP traverse, implemented stack
//Stack define an interface
type Stack interface {
	GetSize() int        // get size of the stack
	IsEmpty() bool       // check if the stack is empty
	Push(element string) // push element to stack
	Pop() string         // fetch the top element of the stack
	Peek() string        // verify the top value
}
type SliceStack []*Node

//NewSliceStack create new SliceStack
func NewSliceStack() SliceStack {
	return make(SliceStack, 0)
}

//Push  an element to SliceStack
func (ss *SliceStack) Push(element *Node) {

	*ss = append(*ss, element)
}

//ToString print the string  list of the SliceStack
func (ss *SliceStack) ToString() string {

	return fmt.Sprintf("Bottom %#v Top", *ss)
}

//Pop will get the pop the top element from SliceStack
func (ss *SliceStack) Pop() (element *Node) {
	length := len(*ss)
	if length == 0 {
		return &Node{}
	}

	ret := (*ss)[length-1]
	if length == 1 {
		*ss = make(SliceStack, 0)
	} else {
		*ss = (*ss)[:length-1]
	}

	return ret
}

//Peek get the top element
func (ss *SliceStack) Peek() (element *Node) {
	return (*ss)[len(*ss)-1]
}

//IsEmpty
func (ss *SliceStack) IsEmpty() bool {
	return len(*ss) == 0
}

func (b *BsTree) PreOrderTraverseNR() {

	if b.root == nil {
		return
	}
	stackBST := NewSliceStack()

	stackBST.Push(b.root)
	for !stackBST.IsEmpty() {

		curNode := stackBST.Pop()
		fmt.Println(curNode.value)

		if curNode.right != nil {
			stackBST.Push(curNode.right)
		}

		if curNode.left != nil {
			stackBST.Push(curNode.left)
		}

	}
}

func TestBstreeBasic(t *testing.T) {

	newBsTree := NewBsTree()

	numbers := []int{33, 21, 77, 99, 36, 15, 76}
	for _, number := range numbers {
		newBsTree.Add(number)
	}

	t.Log("Pre-order traverse: \n")
	newBsTree.PreOrderTraverse()

	t.Log("Pre-order none recurse traverse: \n")
	newBsTree.PreOrderTraverseNR()

	t.Log("In-order traverse: \n")
	newBsTree.InOrderTraverse()

	t.Log("Post-order traverse: \n")
	newBsTree.PostOrderTraverse()
}
