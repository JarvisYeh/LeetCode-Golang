package main

import (
	"container/heap"
	"fmt"
)

type Tuple133 struct {
	i     int
	j     int
	val   int
	index int
}

type MinHeap133 []*Tuple133

func (pq MinHeap133) Len() int { return len(pq) }

func (pq MinHeap133) Less(i, j int) bool { return pq[i].val < pq[j].val }

func (pq MinHeap133) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *MinHeap133) Push(x interface{}) {
	tuple := x.(*Tuple133)
	tuple.index = len(*pq)
	*pq = append(*pq, tuple)
}

func (pq *MinHeap133) Pop() interface{} {
	tuple := (*pq)[len(*pq)-1]
	(*pq)[len(*pq)-1] = nil
	*pq = (*pq)[0 : len(*pq)-1]
	return tuple
}

// merge k sorted array
// use a minHeap to store k pointers
// TC: O(k + knlogk)
// SC: O(k)
func merge133(arrs [][]int) []int {
	// heapify first k pointers
	minHeap := make(MinHeap133, 0)
	idx := 0
	count := 0
	for i := 0; i < len(arrs); i++ {
		if arrs[i] == nil || len(arrs[i]) == 0 {
			continue
		}
		count += len(arrs[i])
		minHeap = append(minHeap, &Tuple133{i, 0, arrs[i][0], idx})
		idx++
	}
	heap.Init(&minHeap)

	res := make([]int, count)
	for i := 0; i < count; i++ {
		curr := heap.Pop(&minHeap).(*Tuple133)
		res[i] = curr.val
		if curr.j+1 < len(arrs[curr.i]) {
			heap.Push(&minHeap, &Tuple133{curr.i, curr.j + 1, arrs[curr.i][curr.j+1], -1})
		}
	}
	return res
}

func main() {
	fmt.Println(
		merge133([][]int{
			{1, 5, 7, 9, 10},
			{-10, 6, 100, 120},
			{55, 57, 59},
			{33, 90, 1000},
		}))
}
