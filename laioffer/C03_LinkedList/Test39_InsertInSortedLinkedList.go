package main

import (
	"fmt"
	"math"
)

// without dummy head
func insert39I(head *ListNode, val int) *ListNode {
	// if head will change
	if head == nil || val < head.Val {
		newHead := &ListNode{val, head}
		return newHead
	}

	// if head won't change
	curr := head
	for curr.Next != nil && curr.Next.Val < val {
		curr = curr.Next
	}
	node := &ListNode{val, curr.Next}
	curr.Next = node
	return head
}

// with dummy head
func insert39II(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{math.MaxInt64, head}
	prev := dummyHead
	// would like to insert after prev
	// insert when prev.val < val < prev.next.val
	// or prev.next = nil
	// else prev = prev.next
	for prev.Next != nil && prev.Next.Val < val {
		prev = prev.Next
	}
	node := &ListNode{val, prev.Next}
	prev.Next = node
	return dummyHead.Next
}

func main() {
	head := CreateListWithArray([]int{1, 3, 5, 7, 9})
	for _, i := range []int{0, 2, 12} {
		head = insert39II(head, i)
		fmt.Println(head)
	}
}
