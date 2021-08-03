package main

func longestOnes(nums []int, k int) int {
	left, right, max := 0, 0, 0
	for right < len(nums) {
		if nums[right] == 1 {
			right++
		} else {
			if k > 0 {
				right++
				k--
			} else {
				if nums[left] == 0 {
					k++
				}
				left++
			}
		}

		// update max
		if right-left > max {
			max = right - left
		}
	}
	return max
}
