package main

import (
	"fmt"
	"sort"
)

// TC: O(n^2 + nlogn)
// SC: O(1) if heap sort
func allTriples186(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		// skip duplicate elements
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// do two sum for (i, end)
		j, k := i+1, len(nums)-1
		for j < k {
			if nums[j]+nums[k] > target {
				k--
			} else if nums[j]+nums[k] < target {
				j++
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				j++
				k--
			}
		}
	}
	return res
}

func main() {
	fmt.Println(allTriples186([]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}, 0))
}
