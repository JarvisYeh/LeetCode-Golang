package main

func swapPairs24(head *ListNode) *ListNode {
	// base case: no node or one node
	if head == nil || head.Next == nil {
		return head
	}

	// 1 -> 2 -> [next ...
	next := swapPairs24(head.Next.Next)
	// 1 -> 2(newHead) -> [next ...
	newHead := head.Next
	// 1 <-> 2(newHead) [next ...
	newHead.Next = head
	// 2 -> 1 -> [next ...
	head.Next = next
	// return 2 to previous level
	return newHead
}
