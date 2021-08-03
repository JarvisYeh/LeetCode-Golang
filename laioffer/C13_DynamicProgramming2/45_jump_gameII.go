package main

func jump45(nums []int) int {
	n := len(nums)
	// minStep[i]: minimal jump steps from i to end
	minSteps := make([]int, len(nums))
	minSteps[n-1] = 0

	for i := n - 2; i >= 0; i-- {
		// initialize m[i] = -1 as it can not jump from i to end
		minSteps[i] = -1
		// check range [i, i + nums[i]], the range that i can jump to
		for j := i; j <= i+nums[i]; j++ {
			// if j in range can jump to end i.e. m[j] != -1
			if j < n && minSteps[j] != -1 {
				// update m[i] if necessary
				// 1. first time update
				// 2. shorter jump steps
				if minSteps[i] == -1 || minSteps[i] > minSteps[j]+1 {
					minSteps[i] = minSteps[j] + 1
				}
			}
		}
	}
	return minSteps[0]
}
