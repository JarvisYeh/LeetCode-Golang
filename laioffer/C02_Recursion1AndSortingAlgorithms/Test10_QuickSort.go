package main

import "math/rand"

// TC: worst case: O(n^2) avg case: O(nlog n)
// SC: worst case: O(n) avg case: O(log n)
func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	// generate pivot index
	randIndex := left + rand.Intn(right-left+1)
	// swap the pivot to the end of the array
	nums[right], nums[randIndex] = nums[randIndex], nums[right]
	// [l, i) < pivot
	// (j, r - 1] > pivot
	// r = pivot
	i, j, pivot := left, right-1, nums[right]
	for i <= j {
		if nums[i] < pivot {
			i++
		} else {
			nums[j], nums[i] = nums[i], nums[j]
			j--
		}
	}
	// nums[i] > pivot by definition
	nums[i], nums[right] = nums[right], nums[i]

	// recursively sort [l, pivot) and (pivot, r]
	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)
}
