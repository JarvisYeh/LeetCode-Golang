package main

func arrayShuffle(nums []int) []int {
	i, j := 0, len(nums)-1
	// [0, i) non-zero
	// [i, j] to be checked
	// (j, len - 1] zeros
	for i <= j {
		if nums[i] == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		} else {
			i++
		}
	}
	return nums
}
