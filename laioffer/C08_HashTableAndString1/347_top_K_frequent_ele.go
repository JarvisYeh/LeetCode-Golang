package main

import "container/heap"

func topKFrequent347(nums []int, k int) []int {
	// use a hashmap to count num frequency
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	// use minHeap to get the top k elements
	pq := &PriorityQueue347{}
	for key, count := range m {
		if len(*pq) < k {
			heap.Push(pq, &Tuple347{key, count})
		} else {
			if (*pq)[0].Count < count {
				heap.Pop(pq)
				heap.Push(pq, &Tuple347{key, count})
			}
		}
	}

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = heap.Pop(pq).(*Tuple347).Key
	}
	return res
}

type Tuple347 struct {
	Key   int
	Count int
}

type PriorityQueue347 []*Tuple347

func (pq PriorityQueue347) Less(i, j int) bool {
	return pq[i].Count < pq[j].Count
}

func (pq PriorityQueue347) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PriorityQueue347) Len() int {
	return len(pq)
}

func (pq *PriorityQueue347) Push(x interface{}) {
	*pq = append(*pq, x.(*Tuple347))
}

func (pq *PriorityQueue347) Pop() interface{} {
	popped := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return popped
}
