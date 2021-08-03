package main

// M[i] represents can jump from i to end
func canJump55I(nums []int) bool {
	M := make([]bool, len(nums))
	M[len(nums)-1] = true
	for i := len(nums) - 2; i >= 0; i-- {
		// see whether the positions that nums[i] can reach is true
		for j := i; j <= i+nums[i]; j++ {
			if j < len(nums) && M[j] {
				M[i] = true
				break
			}
		}
	}
	return M[0]
}

// M[i] represents i can be reached from start
func canJump55II(nums []int) bool {
	M := make([]bool, len(nums))
	M[0] = true
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if M[j] && j+nums[j] >= i {
				M[i] = true
			}
		}
	}
	return M[len(nums)-1]
}
