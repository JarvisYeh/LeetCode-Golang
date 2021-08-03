package main

// sorted array
// remove duplicates for >= 2 elements
func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	// start form 2, since nums[0, 1] can be the same
	i, j := 2, 2 // keep [0, i) check nums[j]
	for j < len(nums) {
		// current element same as previous 2 element (since sorted)
		if nums[j] == nums[i-2] {
			j++
		} else {
			nums[i] = nums[j]
			i++
			j++
		}
	}
	return i
}
