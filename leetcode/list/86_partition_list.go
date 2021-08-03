package main

import "fmt"

func partition86(head *ListNode, x int) *ListNode {
	dummyHead1 := &ListNode{-1, nil}
	dummyHead2 := &ListNode{-1, nil}
	curr1, curr2 := dummyHead1, dummyHead2
	curr := head
	for curr != nil {
		if curr.Val < x {
			curr1.Next = curr
			curr1 = curr1.Next
		} else {
			curr2.Next = curr
			curr2 = curr2.Next
		}
		curr = curr.Next
	}
	curr1.Next = dummyHead2.Next
	curr2.Next = nil
	return dummyHead1.Next
}

func main() {
	head := CreateListWithArray([]int{1, 4, 3, 2, 5, 2})
	fmt.Println(partition86(head, 3))
}
