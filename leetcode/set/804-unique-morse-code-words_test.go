package uniquemorserepresentations

import (
	"fmt"
	"strings"
	"testing"
)

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
	t.Log(uniqueMorseRepresentations(words))
}
