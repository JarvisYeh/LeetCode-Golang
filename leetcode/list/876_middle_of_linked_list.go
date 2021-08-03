package main

import "fmt"

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func main() {
	head := CreateListWithArray([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(middleNode(head).Val)
}
