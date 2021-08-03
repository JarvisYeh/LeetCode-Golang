package main

func maxSubArray53(nums []int) int {
	subSums := make([]int, len(nums))
	subSums[0] = nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if subSums[i-1] < 0 {
			subSums[i] = nums[i]
		} else {
			subSums[i] = subSums[i-1] + nums[i]
		}
		if subSums[i] > max {
			max = subSums[i]
		}
	}
	return max
}
