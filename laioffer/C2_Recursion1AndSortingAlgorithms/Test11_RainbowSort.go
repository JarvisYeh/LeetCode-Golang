package main

func rainbowSort(nums []int) {
	// [0, i) <0
	// [i, j) =0
	// [j, k] to be checked
	// (k, len - 1] >0
	i, j, k := 0, 0, len(nums)-1
	for j <= k {
		if nums[j] < 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		} else if nums[j] == 0 {
			j++
		} else {
			nums[j], nums[k] = nums[k], nums[j]
			k--
		}
	}
}
