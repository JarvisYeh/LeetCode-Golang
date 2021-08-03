package main

// TC: O(k*2^n)
//  for each bucket, find a suitable subset is O(2^n)
//	there are k bucket
// SC: O(n)
func canPartitionKSubsets698(nums []int, k int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%k != 0 {
		return false
	}
	return dfs698(nums, 0, 0, sum/k, k, make([]bool, len(nums)))
}

func dfs698(nums []int, startIndex, currSum, target int, k int, used []bool) bool {
	// k能减少到1说明k>1的时候每个阶段都已经找到sum/k的组合
	// 剩下的值一定等于sum/k,直接return true
	// base case
	if k == 1 {
		return true
	}

	// 对当前k已经找到了和为target的组合，对k-1寻找
	if currSum == target {
		return dfs698(nums, 0, 0, target, k-1, used)
	}

	for i := startIndex; i < len(nums); i++ {
		if !used[i] && currSum+nums[i] <= target {
			used[i] = true
			if dfs698(nums, i+1, currSum+nums[i], target, k, used) {
				return true
			}
			used[i] = false
		}
	}
	return false
}

func main() {
	canPartitionKSubsets698([]int{4, 3, 2, 3, 5, 2, 1}, 4)
}
