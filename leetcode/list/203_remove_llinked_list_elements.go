package main

func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{-1, nil}
	dummyHead.Next = head
	prev := dummyHead
	curr := head
	for curr != nil {
		if curr.Val == val {
			prev.Next = curr.Next
			curr = curr.Next
		} else {
			prev = curr
			curr = curr.Next
		}

	}
	return dummyHead.Next
}
