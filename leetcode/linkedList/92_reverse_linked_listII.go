package main

import "fmt"

func reverseBetween92(head *ListNode, leftIdx int, rightIdx int) *ListNode {
	dummyHead := &ListNode{-1, nil}
	dummyHead.Next = head
	prev, curr := dummyHead, head

	for i := 1; i < leftIdx; i++ {
		prev = curr
		curr = curr.Next
	}

	// reverse [curr, curr + (right - left)]
	subTail := curr
	subHead, next := reverse92(curr, rightIdx-leftIdx)

	// connect subTail and subHead
	prev.Next = subHead
	subTail.Next = next

	return dummyHead.Next
}

func reverse92(curr *ListNode, count int) (*ListNode, *ListNode) {
	var prev *ListNode
	// jump count + 1 steps so that current is jumped out of the sub linked list
	for i := 0; i <= count; i++ {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev, curr
}

func main() {
	head := CreateListWithArray([]int{1, 2, 3, 4, 5})
	fmt.Println(reverseBetween92(head, 2, 4))
}
