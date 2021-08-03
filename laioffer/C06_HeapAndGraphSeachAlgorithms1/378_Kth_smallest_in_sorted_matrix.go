package main

import (
	"container/heap"
	"fmt"
	"strconv"
)

type Tuple378 struct {
	i     int
	j     int
	val   int
	index int
}

// define a priority queue, tuple val as priority

type PriorityQueue378 []*Tuple378

func (pq PriorityQueue378) Len() int {
	return len(pq)
}

func (pq PriorityQueue378) Less(i, j int) bool {
	return pq[i].val < pq[j].val
}

func (pq PriorityQueue378) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = j
	pq[j].index = i
}

func (pq *PriorityQueue378) Push(x interface{}) {
	pushed := x.(*Tuple378)
	pushed.index = len(*pq)
	*pq = append(*pq, pushed)
}

func (pq *PriorityQueue378) Pop() interface{} {
	n := len(*pq)
	popped := (*pq)[n-1]
	popped.index = -1
	(*pq)[n-1] = nil
	*pq = (*pq)[:n-1]
	return popped
}

// for debug, fmt print the pq
func (pq *PriorityQueue378) String() string {
	res := "["
	for _, tuple := range *pq {
		res += strconv.Itoa(tuple.val) + "->"
	}
	return res + "]\n"
}

// solution starts here

// input: n*n matrix
func kthSmallest378(matrix [][]int, k int) int {
	n := len(matrix)
	pq := &PriorityQueue378{}
	heap.Init(pq)
	heap.Push(pq, &Tuple378{0, 0, matrix[0][0], -1})

	// create a all false generated matrix
	generated := make([][]bool, n)
	for i := range generated {
		generated[i] = make([]bool, n)
	}

	for k > 1 {
		popped := heap.Pop(pq).(*Tuple378)
		x, y := popped.i, popped.j
		if x+1 < n && !generated[x+1][y] {
			heap.Push(pq, &Tuple378{x + 1, y, matrix[x+1][y], -1})
			generated[x+1][y] = true
		}
		if y+1 < n && !generated[x][y+1] {
			heap.Push(pq, &Tuple378{x, y + 1, matrix[x][y+1], -1})
			generated[x][y+1] = true
		}
		k--
	}
	return heap.Pop(pq).(*Tuple378).val
}

func main() {
	fmt.Println(kthSmallest378([][]int{{1, 3, 5}, {6, 7, 12}, {11, 14, 14}}, 6))
}
