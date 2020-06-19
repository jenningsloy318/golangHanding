// Given a non-empty, singly linked list with head node head, return a middle node of linked list.
//
// If there are two middle nodes, return the second middle node.
//
//
//
// Example 1:
//
// Input: [1,2,3,4,5]
// Output: Node 3 from this list (Serialization: [3,4,5])
// The returned node has value 3.  (The judge's serialization of this node is [3,4,5]).
// Note that we returned a ListNode object ans, such that:
// ans.val = 3, ans.next.val = 4, ans.next.next.val = 5, and ans.next.next.next = NULL.
// Example 2:
//
// Input: [1,2,3,4,5,6]
// Output: Node 4 from this list (Serialization: [4,5,6])
// Since the list has two middle nodes with values 3 and 4, we return the second one.
//
//
// Note:
//
// The number of nodes in the given list will be between 1 and 100.
//

package MiddleoftheLinkedList

import (
	"fmt"
	"testing"
)

/**
 * Definition for singly-linked list.
**/
type ListNode struct {
	Val  int
	Next *ListNode
}

func addListNode(node *ListNode, val int) *ListNode {

	newNode := &ListNode{
		Val: val,
	}

	if node == nil {
		node = newNode
		return node
	}

	if node.Next != nil {
		node.Next = addListNode(node.Next, val)
	} else {
		node.Next = newNode
	}

	return node

}

func toString(head *ListNode) string {
	var allString []int
	for node := head; node != nil; node = node.Next {
		allString = append(allString, node.Val)

	}
	return fmt.Sprintf("%v\n", allString)
}

func middleNode(head *ListNode) *ListNode {
	var size int
	for curNode := head; curNode != nil; curNode = curNode.Next {
		size++
	}

	newSize := size/2 + 1

	curNode := head
	for i := 1; i < newSize; i++ {
		curNode = curNode.Next
	}

	return curNode
}

func TestConvertBinaryNumberInALinkedListtoInteger(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	var Head *ListNode

	for _, val := range input {
		Head = addListNode(Head, val)
	}
	t.Logf("Original Linked list: %s", toString(Head))

	t.Logf("Middle of the linked list is: %s", toString(middleNode(Head)))

}
