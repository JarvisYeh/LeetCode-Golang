package main

// DP
// TC: O(n^2)
// SC: O(n^2)
func canWin657(nums []int) int {
	// max[i][j] is the max size one can get with range [i, j] pizzas
	max := make([][]int, len(nums))
	for i, _ := range max {
		max[i] = make([]int, len(nums))
	}

	for j := 0; j < len(nums); j++ {
		for i := j; i >= 0; i-- {
			if i == j {
				max[i][j] = nums[i]
			} else if i == j-1 {
				if nums[i] > nums[j] {
					max[i][j] = nums[i]
				} else {
					max[i][j] = nums[j]
				}
			} else {
				left, right := nums[i], nums[j]
				if nums[i+1] > nums[j] {
					left += max[i+2][j]
				} else {
					left += max[i+1][j-1]
				}

				if nums[i] > nums[j-1] {
					right += max[i+1][j-1]
				} else {
					right += max[i][j-2]
				}

				// max[i][j] = max(left, right)
				if left > right {
					max[i][j] = left
				} else {
					max[i][j] = right
				}
			}
		}
	}
	return max[0][len(max)-1]
}
