package main

import (
	"fmt"
)

// TC: O(n^2)
// SC: O(1)
func selectionSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[minIndex] > nums[j] {
				minIndex = j
			}
		}
		nums[minIndex], nums[i] = nums[i], nums[minIndex]
	}
	return nums
}

// selection sort with 3 stacks
// TC: O(n^2)
// SC: O(n)
func selectionSortWithThreeStack(nums []int) []int {
	in := nums
	var buf, res []int // creates nil slices
	// buf := make([]int, 0) or buf := []int{} creates an empty slice
	// res := make([]int, 0) or  res := []int{} creates an empty slice
	for len(in) > 0 {
		min := in[len(in)-1]
		for len(in) > 0 {
			// pop from in stack
			remove := in[len(in)-1]
			in = in[:len(in)-1]
			if remove < min {
				min = remove
			}
			buf = append(buf, remove)
		}

		for len(buf) > 0 {
			// pop from buffer stack
			remove := buf[len(buf)-1]
			buf = buf[:len(buf)-1]
			if remove == min {
				res = append(res, min)
			} else {
				in = append(in, remove)
			}
		}
	}
	return res
}

func main() {
	fmt.Println(selectionSortWithThreeStack([]int{5, 2, 3, 1}))
}
