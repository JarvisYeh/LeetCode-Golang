package main

func findMaxConsecutiveOnes485(nums []int) int {
	count := 0
	max := count
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
			if count > max {
				max = count
			}
		} else {
			count = 0
		}
	}
	return max
}
