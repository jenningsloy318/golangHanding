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

func TestLinkedList(t *testing.T) {

	var head = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val:  6,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}

	t.Logf("Removing %d from %v", 6, toString(head))
	result := removeElements(head, 6)

	t.Logf("After remove %s", toString(result))

}
