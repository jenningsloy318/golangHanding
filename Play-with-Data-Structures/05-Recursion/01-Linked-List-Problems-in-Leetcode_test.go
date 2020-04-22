package linkedlist

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

    
	for  previousNode := dummyHead; previousNode.Next != nil;  {
		if previousNode.Next.Val == val {
			previousNode.Next = previousNode.Next.Next
		} else {
			previousNode = previousNode.Next
		}
	}

	return dummyHead.Next
}
