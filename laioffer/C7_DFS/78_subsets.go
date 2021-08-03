package main

import "fmt"

func subsets78I(nums []int) [][]int {
	res := &[][]int{}
	subset := &[]int{}
	findSubset78(0, nums, subset, res)
	return *res
}

func findSubset78(level int, nums []int, subset *[]int, res *[][]int) {
	if level == len(nums) {
		list := []int{}
		list = append(list, *subset...)
		*res = append(*res, list)
		return
	}

	*subset = append(*subset, nums[level])
	findSubset78(level+1, nums, subset, res)
	*subset = (*subset)[:len(*subset)-1]
	findSubset78(level+1, nums, subset, res)
}

// bit-manipulation method
func subsets78II(nums []int) [][]int {
	res := [][]int{}
	subsetsAmount := 1 << len(nums) // 2^len(nums)

	for i := 0; i < subsetsAmount; i++ { // i = 0b00, 0b01, 0b10, 0b11 for len(nums) = 2
		subset := []int{}
		for j := 0; j < len(nums); j++ {
			if (i >> j & 1) == 1 {
				subset = append(subset, nums[j])
			}
		}
		res = append(res, subset)
	}
	return res
}

func main() {
	fmt.Println(subsets78II([]int{1, 2}))
}
