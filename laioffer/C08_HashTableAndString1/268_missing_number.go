package main

import "sort"

// use xor
// SC: O(1)
// TC: O(n)
// no overflow
func missingNumber268I(nums []int) int {
	num := 0
	for i := 0; i <= len(nums); i++ {
		num ^= i
	}

	for _, n := range nums {
		num ^= n
	}
	return num
}

// use sum
// SC: O(1)
// TC: O(n)
// sum may overflow
func missingNumber268II(nums []int) int {
	sum := (0 + len(nums)) * (len(nums) + 1) / 2
	for _, n := range nums {
		sum -= n
	}
	return sum
}

// use hashmap
// SC: O(n)
// TC: O(n)
func missingNumber268III(nums []int) int {
	m := make(map[int]bool)
	for _, n := range nums {
		m[n] = true
	}

	for i := 0; i <= len(nums); i++ {
		if !m[i] {
			return i
		}
	}
	return -1
}

// binary search
// binary search
// 0 1 2 3
// 0 1 3 4
// find first nums[i] > i, answer is i - 1
func missingNumber268IV(nums []int) int {
	// sort consume most time
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	l, r := 0, len(nums)-1
	for l < r-1 {
		mid := l + (r-l)/2
		if nums[mid] == mid {
			l = mid + 1
		} else if nums[mid] > mid {
			r = mid
		}
	}

	// post-processing
	if nums[l] == l && nums[r] == r { // 0 1 2 3 -> 0 1 2 3 -> l = 2, r = 3, missing 4
		return r + 1
	} else if nums[l] != l {
		return nums[l] - 1
	} else {
		return nums[r] - 1
	}
}
