package main

func permute46I(nums []int) [][]int {
	used := make([]bool, len(nums))
	res := &[][]int{}
	comb := make([]int, len(nums))
	findPermu46(0, nums, &used, &comb, res)
	return *res
}

func findPermu46(index int, nums []int, used *[]bool, comb *[]int, res *[][]int) {
	if index == len(nums) {
		tmp := []int{}
		tmp = append(tmp, *comb...)
		*res = append(*res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			(*comb)[index] = nums[i]
			(*used)[i] = true
			findPermu46(index+1, nums, used, comb, res)
			(*used)[i] = false
		}
	}
}

// use swap
func permute46II(nums []int) [][]int {
	res := &[][]int{}
	findPermu46II(0, nums, res)
	return *res
}

func findPermu46II(index int, nums []int, res *[][]int) {
	if index == len(nums)-1 {
		*res = append(*res, nums)
		return
	}

	// [0, index) is already determined
	// now determine value of nums[index] using swap
	for i := index; i < len(nums); i++ {
		nums[index], nums[i] = nums[i], nums[index]
		findPermu46II(index+1, nums, res)
		nums[index], nums[i] = nums[i], nums[index]
	}
}
