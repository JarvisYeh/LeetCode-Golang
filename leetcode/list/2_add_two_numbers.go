package main

import "fmt"

// add two numbers
// numbers are already reversed order
// return the reversed order
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{-1, nil}
	curr := dummyHead
	overflow := 0

	// when there is carry or any 1 list is not finished
	for l1 != nil || l2 != nil || overflow == 1 {
		sum := overflow
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		overflow = sum / 10
		newNode := &ListNode{sum % 10, nil}
		curr.Next = newNode
		curr = curr.Next
	}

	return dummyHead.Next
}

func main() {
	one := CreateListWithArray([]int{9, 9, 9, 9, 9, 9, 9})
	two := CreateListWithArray([]int{9, 9, 9, 9})
	fmt.Println(addTwoNumbers2(one, two))
}
