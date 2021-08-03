package main

import "fmt"

// SC: O(n^2)
// since every level of recursion tree has a hashset
func permuteUnique47(nums []int) [][]int {
	res := &[][]int{}
	findPermutations47(0, nums, res)
	return *res
}

func findPermutations47(index int, nums []int, res *[][]int) {
	// base case
	if index >= len(nums)-1 {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}

	set := make(map[int]bool)
	for i := index; i < len(nums); i++ {
		if _, ok := set[nums[i]]; !ok {
			set[nums[i]] = true
			nums[index], nums[i] = nums[i], nums[index]
			findPermutations47(index+1, nums, res)
			nums[index], nums[i] = nums[i], nums[index]
		}
	}

}

func main() {
	fmt.Println(permuteUnique47([]int{1, 1}))
}
