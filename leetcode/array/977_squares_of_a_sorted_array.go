package main

import "math"

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	i, j := 0, len(nums)-1
	k := len(res) - 1
	for i <= j {
		if math.Abs(float64(nums[i])) >= math.Abs(float64(nums[j])) {
			res[k] = nums[i] * nums[i]
			i++
		} else {
			res[k] = nums[j] * nums[j]
			j--
		}
		k--
	}
	return res
}
