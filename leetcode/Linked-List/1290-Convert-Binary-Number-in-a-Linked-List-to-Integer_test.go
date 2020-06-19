// Given head which is a reference node to a singly-linked list. The value of each node in the linked list is either 0 or 1. The linked list holds the binary representation of a number.
//
//  Return the decimal value of the number in the linked list.
//
//
// Example 1:
//
//
// Input: head = [1,0,1]
// Output: 5
// Explanation: (101) in base 2 = (5) in base 10
//
//
// Example 2:
//
// Input: head = [0]
// Output: 0
// Example 3:
//
// Input: head = [1]
// Output: 1
// Example 4:
//
// Input: head = [1,0,0,1,0,0,1,1,1,0,0,0,0,0,0]
// Output: 18880
// Example 5:
//
// Input: head = [0,0]
// Output: 0
//
//
// Constraints:
//
// The Linked List is not empty.
// Number of nodes will not exceed 30.
// Each node's value is either 0 or 1.
//

package ConvertBinaryNumberInALinkedListtoInteger

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

func getDecimalValue(head *ListNode) int {
	var l []int
	for curNode := head; curNode != nil; curNode = curNode.Next {
		l = append(l, curNode.Val)
	}
	var sum int
	size := len(l)
	for index, value := range l {
		a := size - 1 - index
		sum = sum + value<<a
	}
	return sum
}

func TestConvertBinaryNumberInALinkedListtoInteger(t *testing.T) {
	input := []int{1, 0, 1, 1, 1, 0, 0}
	var Head *ListNode

	for _, val := range input {
		Head = addListNode(Head, val)
	}
	t.Logf("Original Linked list: %s", toString(Head))

	t.Logf("Converted value of the Linked List: %d", getDecimalValue(Head))

}
