// Problem:
//
// Remove all elements from a linked list of integers that have value val.
//
// Example:
//
// Input:  1->2->6->3->4->5->6, val = 6
// Output: 1->2->3->4->5
//

package RemoveLinkedListElements

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

func toString(head *ListNode) {
	var allString []int
	for node := head; node != nil; node = node.Next {
		allString = append(allString, node.Val)

	}
	fmt.Printf("%v\n", allString)
}
func removeElements(head *ListNode, val int) *ListNode {

	// check linked list is empty
	if head == nil {
		return nil
	}

	// if not empty, recursive remove the val
	result := removeElements(head.Next, val)

	// check if the head has the val, if yes, return head and continue
	if head.Val == val {
		return result
	}

	// if head don't equal to val, assign result to head.Next
	head.Next = result
	return head
}
func TestRemoveLinkedListElements(t *testing.T) {
	//	1->2->6->3->4->5->6
	input := []int{1, 2, 6, 3, 4, 5, 6}
	var dummmyHead = &ListNode{}
	for _, val := range input {
		dummmyHead = addListNode(dummmyHead, val)
	}
	t.Log("Before removal")
	toString(dummmyHead.Next)
	ret := removeElements(dummmyHead.Next, 6)
	t.Log("After removal")
	toString(ret)

}
