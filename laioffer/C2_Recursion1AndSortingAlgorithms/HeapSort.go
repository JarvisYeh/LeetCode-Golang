package main

import "fmt"

func heapSort(nums []int) []int {
	// corner case
	if nums == nil || len(nums) == 0 {
		return nums
	}

	n := len(nums)
	// heapify the max heap, strat from the last non-leaf node
	for i := (n - 1) / 2; i >= 0; i-- {
		percolateDown(nums, n, i)
	}

	// remove the top from the max Heap one at a time
	// swap the last element to top and re-heapify the max heap
	for i := n - 1; i >= 1; i-- {
		// swap the elements. e.g. remove top from the heap
		nums[0], nums[i] = nums[i], nums[0]
		percolateDown(nums, i, 0)
	}
	return nums
}

// iteration
func percolateDown(nums []int, n, index int) {
	for index < n {
		leftChild := 2*index + 1
		rightChild := 2*index + 2
		candidate := index
		if leftChild < n && nums[leftChild] > nums[index] {
			candidate = leftChild
		}
		if rightChild < n && nums[rightChild] > nums[candidate] {
			candidate = rightChild
		}

		if candidate == index {
			break
		} else {
			nums[candidate], nums[index] = nums[index], nums[candidate]
			index = candidate
		}
	}
}

func main() {
	a := heapSort([]int{5, 1, 1, 2, 0, 0})
	fmt.Println(a)
}
