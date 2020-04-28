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

func newListNodes(s []int) *ListNode {

	if s == nil && len(s) == 0 {
		fmt.Errorf("Invalid slice, the slice is empty")
		return nil
	}

	dummyHead := &ListNode{
		Val:  1024,
		Next: &ListNode{},
	}

	for i, curNode := 0, dummyHead.Next; i < len(s)-1; i, curNode = i+1, curNode.Next {
		curNode.Val = s[i]
		curNode.Next = &ListNode{Val: s[i+1]}
	}

	return dummyHead.Next
}

//递归函数，每一次都需要处理基本问题，此处，每次的逻辑都要对head进行处理，包括是否为空，是否 包含对应的值；步骤是先处理一个基本问题（为空的情况），然后构建更小的范围问题的解决方案，构建完成后然后对基本问题进行判断(更小问题有解了，但是除去更小解决方案后，还剩余一个基本元素，对这个基本元素进行解决)
func removeElements(head *ListNode, val int) *ListNode {

	if head == nil {
		return nil
	}

	result := removeElements(head.Next, val)
	if head.Val == val {
		return result
	}

	head.Next = result
	return head

}

func toString(head *ListNode) string {
	var stringSlice []string
	for currNode := head; currNode != nil; currNode = currNode.Next {
		if currNode.Next != nil {
			stringSlice = append(stringSlice, fmt.Sprintf("%d -->", currNode.Val))
		} else {
			stringSlice = append(stringSlice, fmt.Sprintf("%d --> null", currNode.Val))
		}
	}
	return fmt.Sprintf("%v", stringSlice)
}

func toStringRecursion(head *ListNode) string {
	if head == nil {
		return "null"
	}
	result := toStringRecursion(head.Next)
	return fmt.Sprintf("%v --> %v", head.Val, result)
}
func TestConstruct(t *testing.T) {
	s := []int{1, 2, 6, 3, 4, 5, 6}
	x := newListNodes(s)
	t.Log(toString(x))

}

func TestLinkedList(t *testing.T) {
	s := []int{1, 2, 6, 3, 4, 5, 6}
	head := newListNodes(s)

	t.Logf("Removing %d from %v", 6, toString(head))
	t.Logf("Removing %d from [ %v ]", 6, toStringRecursion(head))

	result := removeElements(head, 6)

	t.Logf("After remove %s", toString(result))
	t.Logf("After remove [ %s ]", toStringRecursion(result))

}
