package main

func majorityElement169(nums []int) int {
	candidate, counter := 0, 0
	for _, num := range nums {
		if counter == 0 {
			candidate = num
			counter = 1
		} else if candidate == num {
			counter++
		} else {
			counter--
		}
	}
	return candidate
}
