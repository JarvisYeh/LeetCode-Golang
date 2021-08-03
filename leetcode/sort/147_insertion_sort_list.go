package main

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	dummyHead := ListNode{math.MinInt64, nil}
	dummyHead.Next = head
	next := head
	for next != nil {

	}
	return head
}
