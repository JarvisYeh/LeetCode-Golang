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
