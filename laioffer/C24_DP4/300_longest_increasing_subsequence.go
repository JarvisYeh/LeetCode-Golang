package main

// TC: O(n^2)
// SC: O(n)
func lengthOfLIS300I(nums []int) int {
	DP := make([]int, len(nums))
	max := 1
	for i := 0; i < len(nums); i++ {
		DP[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] && DP[j]+1 > DP[i] {
				DP[i] = DP[j] + 1
			}
		}
		if DP[i] > max {
			max = DP[i]
		}
	}
	return max
}

// with binary search
// TC: O(nlogn)
// SC: O(n)
func lengthOfLIS300II(nums []int) int {
	lowestEnd := []int{}
	// lowestEnd[i] is the smallest end number among all subsequence with length = i + 1
	for i := 0; i < len(nums); i++ {
		// find in lowestEnd array, first number >= nums[i]
		l, r, targetIdx := 0, len(lowestEnd)-1, -1
		for l < r-1 {
			mid := l + (r-l)/2
			if lowestEnd[mid] < nums[i] {
				l = mid + 1
			} else {
				r = mid
			}
		}

		if l < len(lowestEnd) && lowestEnd[l] >= nums[i] {
			targetIdx = l
		} else if r >= 0 && lowestEnd[r] >= nums[i] {
			targetIdx = r
		}

		if targetIdx == -1 {
			lowestEnd = append(lowestEnd, nums[i])
		} else {
			lowestEnd[targetIdx] = nums[i]
		}
	}
	return len(lowestEnd)
}
