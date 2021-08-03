package main

func moveZeroes(nums []int) {
	i, j := 0, 0
	for j < len(nums) {
		if nums[j] == 0 {
			j++
		} else {
			nums[i] = nums[j]
			i++
			j++
		}
	}

	for i < len(nums) {
		nums[i] = 0
		i++
	}
}
