// Write a program to find the node at which the intersection of two singly linked lists begins.
//
// For example, the following two linked lists:
//
//
// begin to intersect at node c1.
//
//
//
// Example 1:
//
//
// Input: intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
// Output: Reference of the node with value = 8
// Input Explanation: The intersected node's value is 8 (note that this must not be 0 if the two lists intersect). From the head of A, it reads as [4,1,8,4,5]. From the head of B, it reads as [5,0,1,8,4,5]. There are 2 nodes before the intersected node in A; There are 3 nodes before the intersected node in B.
//
//
// Example 2:
//
//
// Input: intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
// Output: Reference of the node with value = 2
// Input Explanation: The intersected node's value is 2 (note that this must not be 0 if the two lists intersect). From the head of A, it reads as [0,9,1,2,4]. From the head of B, it reads as [3,2,4]. There are 3 nodes before the intersected node in A; There are 1 node before the intersected node in B.
//
//
// Example 3:
//
//
// Input: intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
// Output: null
// Input Explanation: From the head of A, it reads as [2,6,4]. From the head of B, it reads as [1,5]. Since the two lists do not intersect, intersectVal must be 0, while skipA and skipB can be arbitrary values.
// Explanation: The two lists do not intersect, so return null.
//
//
// Notes:
//
// If the two linked lists have no intersection at all, return null.
// The linked lists must retain their original structure after the function returns.
// You may assume there are no cycles anywhere in the entire linked structure.
// Your code should preferably run in O(n) time and use only O(1) memory.
//

package IntersectionofTwoLinkedLists

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
	var nodeSlice []int
	for node := head; node != nil; node = node.Next {
		nodeSlice = append(nodeSlice, node.Val)

	}
	return fmt.Sprintf("%v\n", nodeSlice)
}

func find(head, node *ListNode) *ListNode {

	if head == nil || node == nil {
		return nil
	}

	for curNode := head; curNode != nil; curNode = curNode.Next {
		if curNode == node {
			return curNode
		}
	}
	return nil
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	for curNode := headB; curNode != nil; curNode = curNode.Next {
		ret := find(headA, curNode)
		if ret != nil {
			return ret
		}
	}
	return nil
}
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	for curNodeA := headA; curNodeA != nil; curNodeA = curNodeA.Next {

		for curNodeB := headB; curNodeB != nil; curNodeB = curNodeB.Next {
			if curNodeA == curNodeB {
				return curNodeA
			}
		}

	}
	return nil
}

func getIntersectionNode3(headA, headB *ListNode) *ListNode {

	if headA == nil || headB == nil {
		return nil
	}

	var nodeMap = make(map[*ListNode]int)
	for curNode, i := headA, 0; curNode != nil && i <= 1; {
		if _, ok := nodeMap[curNode]; ok {
			return curNode
		} else {
			nodeMap[curNode] = 1
		}
		if curNode.Next == nil {
			curNode = headB
			i++
		} else {
			curNode = curNode.Next
		}

	}

	return nil
}

func TestIntersectionofTwoLinkedLists(t *testing.T) {

	var Head1 *ListNode
	input1 := []int{-10, -10, -9, -4, 1, 6, 6}
	for _, val := range input1 {

		Head1 = addListNode(Head1, val)
	}
	var Head2 *ListNode
	input2 := []int{-7, 1, 3, 4}
	for _, val := range input2 {

		Head2 = addListNode(Head2, val)
	}
	t.Logf("intersection node: %v", getIntersectionNode3(Head1, Head2))

	Head2.Next.Next = Head1.Next.Next
	t.Logf("intersection node: %v", getIntersectionNode3(Head1, Head2))
}
