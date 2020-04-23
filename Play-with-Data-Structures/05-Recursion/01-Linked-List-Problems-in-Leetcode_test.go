package linkedlist

import (
	"fmt"
	"testing"
)

//https://leetcode.com/problems/remove-linked-list-elements/

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func newListNodes(s []int) *ListNode {

	if s == nil && len(s) == 0 {
		fmt.Errorf("Invalid slice, the slice is empty")
		return nil
	}

	dummyHead := &ListNode{
		Val:  1024,
		Next: &ListNode{},
	}

	for i, curNode := 0, dummyHead.Next; i < len(s)-1; i, curNode = i+1, curNode.Next {
		curNode.Val = s[i]
		curNode.Next = &ListNode{Val: s[i+1]}
	}

	return dummyHead.Next
}
func removeElements(head *ListNode, val int) *ListNode {
	var dummyHead = &ListNode{
		Val:  1024,
		Next: head,
	}

	for previousNode := dummyHead; previousNode.Next != nil; {
		if previousNode.Next.Val == val {
			previousNode.Next = previousNode.Next.Next
		} else {
			previousNode = previousNode.Next
		}
	}

	return dummyHead.Next
}

func toString(head *ListNode) string {
	var stringSlice []string
	for currNode := head; currNode != nil; currNode = currNode.Next {
		stringSlice = append(stringSlice, fmt.Sprintf("%d ->", currNode.Val))
	}
	return fmt.Sprintf("%v", stringSlice)

}

func TestConstruct(t *testing.T) {
	s := []int{1, 2, 6, 3, 4, 5, 6}
	x := newListNodes(s)
	t.Log(toString(x))

}

func TestLinkedList(t *testing.T) {
	s := []int{1, 2, 6, 3, 4, 5, 6}
	head := newListNodes(s)

	t.Logf("Removing %d from %v", 6, toString(head))
	result := removeElements(head, 6)

	t.Logf("After remove %s", toString(result))

}
