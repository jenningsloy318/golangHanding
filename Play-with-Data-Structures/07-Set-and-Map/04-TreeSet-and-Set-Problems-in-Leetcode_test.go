package set2

import (
	"fmt"
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

func (b *BsTree) getSize() int {
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

func uniqueMorseRepresentations(words []string) int {
	codes := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	var wordCodesList []string

	for _, word := range words {
		var wordCode string
		var wordCodes []string
		for _, char := range word {
			wordCodes = append(wordCodes, codes[char-'a'])
		}
		wordCode = fmt.Sprintf(strings.Join(wordCodes, ""))
		wordCodesList = append(wordCodesList, wordCode)
	}

	newTree := NewBsTree()
	for _, node := range wordCodesList {
		newTree.Add(node)
	}

	return newTree.size
}

func TestUniqueMorseRepresentations(t *testing.T) {
	words := []string{"gin", "zen", "gig", "msg"}
	count := uniqueMorseRepresentations(words)
	t.Log(count)
}
