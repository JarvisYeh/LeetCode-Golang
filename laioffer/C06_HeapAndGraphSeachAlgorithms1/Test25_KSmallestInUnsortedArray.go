package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"sort"
)

// TC: O(k + (n-k)log k)
// SC: O(K)
// Use maxHeap
func kSmallestI(arr []int, k int) []int {
	// corner case
	if k > len(arr) {
		return arr
	}

	maxHeap := &maxIntHeap25{}
	*maxHeap = append(*maxHeap, arr[0:k]...)
	heap.Init(maxHeap)

	for i := k; i < len(arr); i++ {
		if maxHeap.Len() < k {
			heap.Push(maxHeap, arr[i])
		} else {
			if (*maxHeap)[0] > arr[i] {
				heap.Pop(maxHeap)
				heap.Push(maxHeap, arr[i])
			}
		}
	}
	sort.Slice(*maxHeap, maxHeap.Less)
	return *maxHeap
}

// TC: O(n + klog n)
// SC: O(n)
// Use minHeap
func kSmallestII(arr []int, k int) []int {
	// corner case
	if k > len(arr) {
		return arr
	}

	minHeap := &minIntHeap25{}
	*minHeap = append(*minHeap, arr...)
	heap.Init(minHeap)

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = heap.Pop(minHeap).(int)
	}

	sort.Slice(res, func(i, j int) bool {
		return i < j
	})
	return res
}

// Use quick select
func kSmallestIII(arr []int, k int) []int {
	if k >= len(arr) {
		return arr
	}
	res := make([]int, len(arr))
	copy(res, arr)
	quickSelect(res, 0, len(arr)-1, k)
	return res[:k]
}

func quickSelect(arr []int, left, right, k int) {
	randIndex := left + rand.Intn(right-left+1)
	// swap pivot to end
	arr[randIndex], arr[right] = arr[right], arr[randIndex]

	// search space [i, j]
	// [0, i) < pivot
	// (j, end] > pivot
	i, j := 0, right-1
	for i <= j {
		if arr[i] < arr[right] {
			i++
		} else {
			arr[i], arr[j] = arr[j], arr[i]
			j--
		}
	}
	// swap back pivot
	arr[i], arr[right] = arr[right], arr[i]

	// shrink search range or return if k-th smallest is found
	if i+1 > k {
		quickSelect(arr, left, i, k)
	} else if i+1 < k {
		quickSelect(arr, i+1, right, k)
	} else {
		return
	}
}

func main() {
	arr := []int{-5, 9, 102, 1, 5, 26, 7, 11, 109, 23}
	fmt.Println(kSmallestI(arr, 5))
	fmt.Println(kSmallestII(arr, 5))
	fmt.Println(kSmallestIII(arr, 5))
}

type minIntHeap25 []int

/**
heap.Interface contains
	sort.Interface contains
		Len() int, Less(i, j int) bool, Swap(i, j int)
	Push(x interface{})
	Pop() interface{}
So one custom heap has to implement those methods
**/

func (h minIntHeap25) Len() int {
	return len(h)
}

// Less description
// if Less(i, j) and Less(j, i) both returns false they are consider equal
func (h minIntHeap25) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minIntHeap25) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minIntHeap25) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minIntHeap25) Pop() interface{} {
	old := *h
	popped := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return popped
}

type maxIntHeap25 []int

func (h maxIntHeap25) Len() int {
	return len(h)
}

// Less description
// if Less(i, j) and Less(j, i) both returns false they are consider equal
func (h maxIntHeap25) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h maxIntHeap25) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxIntHeap25) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxIntHeap25) Pop() interface{} {
	old := *h
	popped := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return popped
}
