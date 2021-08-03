package main

import (
	"fmt"
	"sort"
)

// unsorted array
// return all pairs without duplication
// solution 1 : use count map to record frequency in [0, i) and avoid adding duplicate pairs
// TC : O(n)
// SC : O(n) for map
func allPairs182I(nums []int, target int) [][]int {
	// count map counts the frequency of number in range [0, i)
	countMap := make(map[int]int)
	pairs := [][]int{}
	for i := 0; i < len(nums); i++ {
		if _, findNum := countMap[nums[i]]; findNum { // nums[i] has already occurred in [0, i)
			if countMap[nums[i]] == 1 && nums[i]*2 == target {
				pairs = append(pairs, []int{nums[i], nums[i]})
			}
			countMap[nums[i]]++
		} else {
			// if it's first time see nums[i] and found target - nums[i] in map
			if _, findPair := countMap[target-nums[i]]; findPair {
				// find one result pair
				pairs = append(pairs, []int{target - nums[i], nums[i]})
			}
			// update countMap
			countMap[nums[i]] = 1
		}
	}
	return pairs
}

// solution 2:
// sort first and use two pointers, skip same values to avoid duplicate pairs
// TC : O(nlogn)
// SC : O(1) if use heap sort, actually O(n)
func allPairs182II(nums []int, target int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	pairs := [][]int{}
	// init two pointers
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]+nums[j] > target {
			j--
		} else if nums[i]+nums[j] < target {
			i++
		} else { // nums[i] + nums[j] = target
			pairs = append(pairs, []int{nums[i], nums[j]})
			for i < j && nums[i] == nums[i+1] {
				i++
			}
			for i < j && nums[j] == nums[j-1] {
				j--
			}
			// now nums[i] != nums[i + 1] && , so still need to move one step foward
			// vice versa for j
			i++
			j--
		}
	}
	return pairs
}

func main() {
	fmt.Println(allPairs182II([]int{0, -5, 5, -2, 2, 2, 2, -2, 1, 0}, 0))
}
