package bstmap

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

type Map interface {
	Add(key string, value int)
	Remove(key string) int
	Contains(key string) bool
	Get(key string) int
	Set(key string, value int)
	GetSize() int
	IsEmpty() bool
}

type Node struct {
	Key   string
	Value int
	Left  *Node
	Right *Node
}

func NewNode(key string, value int) *Node {
	return &Node{
		Key:   key,
		Value: value,
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

func (b *BsTree) GetSize() int {
	return b.size
}

func (b *BsTree) IsEmpty() bool {
	return b.size == 0
}

func (b *BsTree) Add(key string, value int) {

	b.root = b.add(b.root, key, value)
}

func (b *BsTree) add(node *Node, key string, value int) *Node {

	if node == nil {
		b.size++
		return NewNode(key, value)
	}

	if key < node.Key {
		node.Left = b.add(node.Left, key, value)
	} else if key > node.Key {
		node.Right = b.add(node.Right, key, value)
	} else {
		node.Value = value
	}

	return node
}

// 返回以node为根节点的二分搜索树中，key所在的节点
func getNode(node *Node, key string) *Node {
	if node == nil {
		return nil
	}

	if key < node.Key {
		return getNode(node.Left, key)
	}
	if key > node.Key {
		return getNode(node.Right, key)
	}
	return node
}

func (b *BsTree) Contains(key string) bool {
	node := getNode(b.root, key)

	if node == nil {
		return false
	} else {
		return true
	}

}

func (b *BsTree) Get(key string) int {
	node := getNode(b.root, key)
	if node == nil {
		return -1
	}
	return node.Value

}

func (b *BsTree) Set(key string, value int) {
	node := getNode(b.root, key)
	if node != nil {
		node.Value = value
	} else {
		fmt.Errorf("Key %s doesn't exist", key)
	}
}

// 返回以node为根的二分搜索树的最小值所在的节点

func minNode(node *Node) *Node {
	if node.Left == nil {
		return node
	} else {
		return minNode(node.Left)
	}
}

// 删除掉以node为根的二分搜索树中的最小节点
// 返回删除节点后新的二分搜索树的根

func (b *BsTree) removeMinNode(node *Node) *Node {
	if node.Left == nil {
		nodeRight := node.Right
		node.Right = nil
		b.size--
		return nodeRight
	} else {
		node.Left = b.removeMinNode(node.Left)
		return node
	}
}
func (b *BsTree) Remove(key string) int {
	node := getNode(b.root, key)

	if node != nil {
		b.root = b.remove(b.root, key)
		return node.Value
	} else {
		return -1
	}

}
func (b *BsTree) remove(node *Node, key string) *Node {

	if key < node.Key {
		return b.remove(node.Left, key)
	} else if key > node.Key {
		return b.remove(node.Right, key)
	} else { // key==node.key
		if node.Left == nil {
			nodeRight := node.Right
			node.Right = nil
			b.size--
			return nodeRight
		}
		if node.Right == nil {
			nodeLeft := node.Left
			node.Left = nil
			b.size--
			return nodeLeft
		}

		successorNode := minNode(node.Right)
		successorNode.Right = b.removeMinNode(node.Right)
		successorNode.Left = node.Left
		node.Left = nil
		node.Right = nil
		return successorNode
	}
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

	newMap := NewBsTree()
	t1 := time.Now()
	for _, word := range testSlice {

		if newMap.Contains(word) {
			newMap.Set(word, newMap.Get(word)+1)
			//t.Log("yes", word, newMap.Get(word))
		} else {
			newMap.Add(word, 1)
			//t.Log("no", word, newMap.Get(word))
		}

	}
	duration := time.Since(t1)
	t.Logf("Total uniq Words: %d, process time %fs ", newMap.GetSize(), duration.Seconds())
	t.Logf("Frequence of  to is : %d.\n", newMap.Get("to"))
	t.Log("Remove word to .")
	newMap.Remove("to")
	t.Logf("Now the frequence of  to is : %d.\n", newMap.Get("to"))

	intMap := make(map[string]int)
	t2 := time.Now()
	for _, word := range testSlice {

		if _, ok := intMap[word]; ok {
			intMap[word] += 1
		} else {
			intMap[word] = 1
		}

	}
	duration2 := time.Since(t2)
	t.Logf("Total uniq Words: %d, process time %fs ", len(intMap), duration2.Seconds())
	t.Logf("Frequence of  to is : %d.\n", intMap["to"])
	delete(intMap, "to")
	t.Logf("Frequence of  to is : %d.\n", intMap["to"])

}
