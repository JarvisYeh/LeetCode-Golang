package main

import "container/heap"

type ListNode23 struct {
	Val  int
	Next *ListNode23
}

type Node23 struct {
	listnode *ListNode23
	index    int
}

type MinHeap23 []*Node23

func (pq MinHeap23) Len() int { return len(pq) }

func (pq MinHeap23) Less(i, j int) bool { return pq[i].listnode.Val < pq[j].listnode.Val }

func (pq MinHeap23) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *MinHeap23) Push(x interface{}) {
	x.(*Node23).index = len(*pq)
	(*pq) = append(*pq, x.(*Node23))
}

func (pq *MinHeap23) Pop() interface{} {
	node := (*pq)[len(*pq)-1]
	(*pq)[len(*pq)-1] = nil
	*pq = (*pq)[:len(*pq)-1]
	return node
}

// method 1: use iteration
// TC: O(k + kn*log k)
// SC: O(k)
func mergeKLists23I(lists []*ListNode23) *ListNode23 {
	// heapify all valid listNode in lists w.r.t Val
	// O(k)
	minHeap := &MinHeap23{}
	for _, listnode := range lists {
		if listnode == nil {
			continue
		}
		*minHeap = append(*minHeap, &Node23{listnode, len(*minHeap)})
	}
	heap.Init(minHeap)

	dummyHead := &ListNode23{-1, nil}
	prev := dummyHead
	// O(kn)
	for len(*minHeap) != 0 {
		curr := heap.Pop(minHeap).(*Node23).listnode //O(log k)
		if curr.Next != nil {
			heap.Push(minHeap, &Node23{curr.Next, -1}) //O(log k)
		}
		prev.Next = curr
		prev = prev.Next
	}
	return dummyHead.Next
}

// method 2: recursion
// TC: every level merge kn listNodes
//		O(kn*log K)
// SC: O(log K)
func mergeKLists(lists []*ListNode23) *ListNode23 {
	if lists == nil || len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}
	mid := len(lists) / 2
	size := len(lists)
	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:size])
	return mergeTwoList(left, right)
}

func mergeTwoList(l1, l2 *ListNode23) *ListNode23 {
	dummyHead := &ListNode23{-1, nil}
	prev := dummyHead
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}
	if l1 != nil {
		prev.Next = l1
	}
	if l2 != nil {
		prev.Next = l2
	}
	return dummyHead.Next
}
