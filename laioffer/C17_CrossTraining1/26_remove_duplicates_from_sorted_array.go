package main

// sorted
// i, j move same direction, relatively order not change
// do this in place, return new length
func removeDuplicates26(nums []int) int {
	i, j := 0, 0
	for j < len(nums) {
		if i == 0 || nums[i-1] != nums[j] {
			nums[i] = nums[j]
			i++
			j++
		} else {
			j++
		}
	}
	return i
}
