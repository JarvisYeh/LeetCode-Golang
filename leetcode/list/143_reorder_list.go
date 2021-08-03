package main

func reorderList143(head *ListNode) {
	mid := findMiddle(head)
	one, two := head, mid.Next

	// split two list
	mid.Next = nil
	two = reverseList(two)

	dummyHead := &ListNode{-1, nil}
	curr := dummyHead
	for one != nil && two != nil {
		curr.Next = one
		one = one.Next
		curr.Next.Next = two
		two = two.Next
		curr = curr.Next.Next
	}

	if one != nil {
		curr.Next = one
	}
	if two != nil {
		curr.Next = two
	}
	head = dummyHead.Next
}
