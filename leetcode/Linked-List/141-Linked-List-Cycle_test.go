// Given a linked list, determine if it has a cycle in it.
//
// To represent a cycle in the given linked list, we use an integer pos which represents the position (0-indexed) in the linked list where tail connects to. If pos is -1, then there is no cycle in the linked list.
//
//
//
// Example 1:
//
// Input: head = [3,2,0,-4], pos = 1
// Output: true
// Explanation: There is a cycle in the linked list, where tail connects to the second node.
//
//
// Example 2:
//
// Input: head = [1,2], pos = 0
// Output: true
// Explanation: There is a cycle in the linked list, where tail connects to the first node.
//
//
// Example 3:
//
// Input: head = [1], pos = -1
// Output: false
// Explanation: There is no cycle in the linked list.
//
//
//
//
// Follow up:
//
// Can you solve it using O(1) (i.e. constant) memory?

package LinkedListCycle

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
func toString(head *ListNode) string {
	var allString []int
	for node := head; node != nil; node = node.Next {
		allString = append(allString, node.Val)

	}
	return fmt.Sprintf("%v\n", allString)
}

func makeCycle(head *ListNode, index int) *ListNode {

	tailNode := head
	for tailNode.Next != nil {
		tailNode = tailNode.Next
	}

	indexNode := head

	for loc, indexNode := 0, head; loc <= index; loc++ {
		indexNode = indexNode.Next
	}

	tailNode.Next = indexNode
	return head
}

func hasCycle(head *ListNode) bool {

	var addrSum = make(map[*ListNode]int)

	for curNode := head; ; curNode = curNode.Next {

		if curNode == nil {
			return false
		} else {
			_, ok := addrSum[curNode]
			if ok {
				return true
			} else {
				addrSum[curNode] = 1
			}
		}
	}
}

//hasCycle2
/*
onsidering two pointers at different speed - a slow pointer and a fast pointer.
The slow pointer moves one step at a time while the fast pointer moves two steps at a time.

If there is no cycle in the list, the fast pointer will eventually reach the end and we can return false in this case.

Now consider a cyclic list and imagine the slow and fast pointers are two runners racing around a circle track.
The fast runner will eventually meet the slow runner.
*/
func hasCycle2(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return false
	}

	slowPointer := head
	fastPointer := head.Next

	for slowPointer != fastPointer {
		if fastPointer == nil || fastPointer.Next == nil {
			return false
		}

		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next.Next
	}

	return true
}

func TestLinkedListCycle(t *testing.T) {
	//	1->2->6->3->4->5->6
	input := []int{1, 2, 6, 3, 4, 5, 6}
	var Head *ListNode
	for _, val := range input {
		Head = addListNode(Head, val)
	}
	t.Logf("Original Linked list: %v", toString(Head))

	t.Logf("Linked list is cycled: %t", hasCycle2(Head))

	t.Logf("Linked list is cycled: %t\n", hasCycle2(makeCycle(Head, 0)))

}
