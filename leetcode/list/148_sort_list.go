package main

func sortList148(head *ListNode) *ListNode {
	return mergeSort148(head)
}

func mergeSort148(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mid := findMiddle(head)
	two := mid.Next
	mid.Next = nil

	return merge148(mergeSort148(head), mergeSort148(two))
}

func merge148(one, two *ListNode) *ListNode {
	dummyHead := &ListNode{-1, nil}
	curr := dummyHead
	for one != nil && two != nil {
		if one.Val < two.Val {
			curr.Next = one
			one = one.Next
		} else {
			curr.Next = two
			two = two.Next
		}
		curr = curr.Next
	}
	if one != nil {
		curr.Next = one
	}
	if two != nil {
		curr.Next = two
	}

	return dummyHead.Next
}
