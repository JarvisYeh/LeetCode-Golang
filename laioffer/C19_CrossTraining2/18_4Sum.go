package main

import "sort"

// 4 sum find all pairs without duplication
func fourSum(nums []int, target int) [][]int {
	// sort first
	sort.Ints(nums)
	// do general k sum
	return kSum(nums, 0, len(nums)-1, target, 4)
}

func kSum(nums []int, left, right, target, k int) [][]int {
	if k == 2 {
		twoSumPairs := [][]int{}
		for left < right {
			if nums[left]+nums[right] < target {
				left++
			} else if nums[left]+nums[right] > target {
				right--
			} else {
				twoSumPairs = append(twoSumPairs, []int{left, right})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
		return twoSumPairs
	} else {
		res := [][]int{}
		for i := left; i <= right; i++ {
			if i != left && nums[i] == nums[i-1] {
				continue
			} // skip duplicates
			for _, comb := range kSum(nums, i+1, right, target-nums[i], k-1) {
				comb = append(comb, nums[i])
				res = append(res, comb)
			}
		}
		return res
	}
}

func main() {
	fourSum([]int{1, 0, -1, 0, -2, 2}, 0)
}
