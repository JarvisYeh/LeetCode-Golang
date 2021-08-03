package main

import "fmt"

// unsorted array
// find all index pairs <i, j> where nums[i] + nums[j] = target

func allPairs181(nums []int, target int) [][]int {
	indexMap := make(map[int][]int)
	resPairs := [][]int{}
	for i := 0; i < len(nums); i++ {
		// check if {target - nums[i]} in indexMap
		if l, ok := indexMap[target-nums[i]]; ok {
			for _, idx := range l {
				resPairs = append(resPairs, []int{idx, i})
			}
		}

		// update indexMap w.r.t. nums[i]
		if l, ok := indexMap[nums[i]]; ok {
			indexMap[nums[i]] = append(l, i)
		} else {
			indexMap[nums[i]] = []int{i}
		}
	}
	return resPairs
}

func main() {
	fmt.Println(allPairs181([]int{1, 1, 5, 4, 2, -1, 9}, 6))
}
