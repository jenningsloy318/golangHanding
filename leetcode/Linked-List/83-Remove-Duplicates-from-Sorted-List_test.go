//Given a sorted linked list, delete all duplicates such that each element appear only once.
//
//Example 1:
//
//Input: 1->1->2
//Output: 1->2
//Example 2:
//
//Input: 1->1->2->3->3
//Output: 1->2->3

package RemoveDuplicatesfromSortedList

import (
	"fmt"
	"testing"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addListNode(node *ListNode, val int) *ListNode {

	newNode := &ListNode{Val: val}
	if node == nil {
		node = newNode
		return node
	}

	if node.Val < val {
		node.Next = addListNode(node.Next, val)
		return node
	} else {
		newNode.Next = node
		return newNode
	}

}

func toString(head *ListNode) string {
	var allString []int
	for node := head; node != nil; node = node.Next {
		allString = append(allString, node.Val)

	}
	return fmt.Sprintf("%v\n", allString)
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	curNode := deleteDuplicates(head.Next)

	if head.Val == curNode.Val {
		return curNode
	}
	head.Next = curNode
	return head
}
func TestRemoveDuplicatesfromSortedList(t *testing.T) {
	//	1->2->6->3->4->5->6
	input := []int{1, 2, 6, 3, 4, 5, 6, 4, 4, 3, 3}
	var Head *ListNode
	for _, val := range input {
		Head = addListNode(Head, val)
	}
	t.Logf("Before removal: %s", toString(Head))

	t.Logf("After removal: %s", toString(deleteDuplicates(Head)))

}
