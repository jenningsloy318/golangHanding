//Merge two sorted linked lists and return it as a new sorted list. The new list should be made by splicing together the nodes of the first two lists.
//
//Example:
//
//Input: 1->2->4, 1->3->4
//Output: 1->1->2->3->4->4

package ReverseLinkedList

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

// add nodes into list in sorted
func add(head *ListNode, val int) *ListNode {
	newNode := &ListNode{Val: val}
	if head == nil {
		head = newNode
		return head
	}

	if head.Val < val {
		head.Next = add(head.Next, val)
		return head
	} else {
		newNode.Next = head
		return newNode
	}

}
func toString(head *ListNode) {
	var nodeSlice []int
	for node := head; node != nil; node = node.Next {
		nodeSlice = append(nodeSlice, node.Val)

	}
	fmt.Printf("Items in the Linked List: %v\n", nodeSlice)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	var ret = l1
	for curNode := l2; curNode != nil; curNode = curNode.Next {
		ret = add(ret, curNode.Val)
	}
	return ret
}

func TestReverseLinkedList(t *testing.T) {

	var Head1 *ListNode
	input1 := []int{-10, -10, -9, -4, 1, 6, 6}
	for _, val := range input1 {
		Head1 = add(Head1, val)
	}

	toString(Head1)

	var Head2 *ListNode
	input2 := []int{-7}
	for _, val := range input2 {
		Head2 = add(Head2, val)
	}
	toString(Head2)

	toString(mergeTwoLists(Head1, Head2))
}
