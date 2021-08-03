package main

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateListWithArray(arr []int) *ListNode {
	dummyHead := &ListNode{-1, nil}
	curr := dummyHead
	for _, i := range arr {
		newNode := &ListNode{i, nil}
		curr.Next = newNode
		curr = curr.Next
	}
	return dummyHead.Next
}

func (head *ListNode) String() string {
	if head == nil {
		return "empty list"
	}
	var sb strings.Builder
	for head.Next != nil {
		sb.WriteString(fmt.Sprintf("%v->", head.Val))
		head = head.Next
	}
	sb.WriteString(fmt.Sprintf("%v", head.Val))
	return sb.String()
}

/**
Find middle point of a linked list
even number of nodes: return the previous one
*/
func findMiddle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

/**
reverse a linked list and return the new head
*/
func reverseList(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
