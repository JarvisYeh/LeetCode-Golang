package main

func searchRange(nums []int, target int) []int {
	start, end := -1, -1

	if len(nums) == 0 {
		return []int{start, end}
	}

	// find start
	low, high := 0, len(nums)-1
	for low < high-1 {
		mid := low + (high-low)/2
		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}
	// the [low, high] contains only 2 elements, the leftmost target index is start
	if nums[low] == target {
		start = low
	} else if nums[high] == target {
		start = high
	}

	// find end
	low = 0
	high = len(nums) - 1
	for low < high-1 {
		mid := low + (high-low)/2
		if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid
		}
	}
	// the [low, high] contains only 2 elements, the leftmost target index is start
	if nums[high] == target {
		end = high
	} else if nums[low] == target {
		end = low
	}

	return []int{start, end}

}
