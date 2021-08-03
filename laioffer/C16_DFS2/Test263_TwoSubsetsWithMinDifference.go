package main

import (
	"fmt"
	"math"
)

// TC: O(2^n)
// SC: O(n)
func minDifference263(nums []int) int {
	min, sum := new(int), 0
	*min = math.MaxInt64
	for _, val := range nums {
		sum += val
	}
	dfs263(nums, 0, 0, 0, sum, min)
	return *min
}

func dfs263(nums []int, index, added, currSum, sum int, min *int) {
	// base case 1: already added len/2 numbers
	if added == len(nums)/2 {
		if int(math.Abs(float64(currSum)*2-float64(sum))) < *min {
			*min = int(math.Abs(float64(currSum)*2 - float64(sum)))
		}
		return
	}
	// base case 2: index reach end
	if index == len(nums) {
		return
	}

	// add nums[index] to subset
	dfs263(nums, index+1, added+1, currSum+nums[index], sum, min)
	// not add nums[index] to subset
	dfs263(nums, index+1, added, currSum, sum, min)
}

func main() {
	fmt.Println(minDifference263([]int{1, 4, 2, 3, 3}))
}
