// Reverse a singly linked list.
//
// Example:
//
// Input: 1->2->3->4->5->NULL
// Output: 5->4->3->2->1->NULL
// Follow up:
//
// A linked list can be reversed either iteratively or recursively. Could you implement both?
//
//

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

type LinkedList struct {
	dummyHead *ListNode
	size      int
}

func (l *LinkedList) addListNode(node *ListNode, val int) *ListNode {

	newNode := &ListNode{
		Val: val,
	}

	if node == nil {
		node = newNode
		l.size++
		return node
	}

	if node.Next != nil {
		node.Next = l.addListNode(node.Next, val)
	} else {
		node.Next = newNode
		l.size++
	}

	return node

}

func toString(head *ListNode) {
	var allString []int
	for node := head; node != nil; node = node.Next {
		allString = append(allString, node.Val)

	}
	fmt.Printf("Items in the Linked List: %v\n", allString)
}

func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	curNode := reverseList(head.Next)

	head.Next.Next = head // point Next node's Next  back to to its previous node
	head.Next = nil       // make this node disconnect from the Linked list as it is being reversing
	return curNode
}

func reverseList2(head *ListNode) *ListNode {
	var prevNode *ListNode
	for currNode := head; currNode != nil; {
		nextTemp := currNode.Next
		currNode.Next = prevNode
		prevNode = currNode
		currNode = nextTemp
	}
	return prevNode
}

func TestReverseLinkedList(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	var dummyHead = &ListNode{}

	l := &LinkedList{
		dummyHead: dummyHead,
		size:      0,
	}

	for _, val := range input {
		l.addListNode(dummyHead, val)
	}
	toString(l.dummyHead.Next)
	t.Logf("Input size: %d", l.size)

	toString(reverseList(l.dummyHead.Next))
}
