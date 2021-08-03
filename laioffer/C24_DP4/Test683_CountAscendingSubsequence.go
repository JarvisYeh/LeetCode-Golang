package main

import (
	"fmt"
)

func numIncreasingSubsequences683(nums []int) int {
	// DP[i] is the number of ascending subsequence ending at i
	DP := make([]int, len(nums))
	sum := 0
	for i := 0; i < len(nums); i++ {
		DP[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				DP[i] += DP[j]
			}
		}
		sum += DP[i]
	}
	return sum
}

func main() {
	fmt.Println(numIncreasingSubsequences683([]int{1, 3, 2, 2}))
}
