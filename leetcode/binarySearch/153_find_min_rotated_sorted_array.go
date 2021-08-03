package main

import "fmt"

func findMin153(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + ((high - low) >> 1)
		if nums[mid] > nums[high] { // [low mid | high] or [low(mid) | high]
			low = mid + 1
		} else if nums[mid] < nums[high] { // [low | mid high] or [low mid high]
			high = mid
		}
	}
	// after loop, lwo = high
	return nums[low]
}

func main() {
	fmt.Println(findMin153([]int{4, 5, 7, 0, 1, 2}))
}
