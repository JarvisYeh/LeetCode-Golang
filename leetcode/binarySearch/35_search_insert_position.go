package main

// range最后有值，还需要post processing
func searchInsertI(nums []int, target int) int {
	low, high := 0, len(nums)-1

	// find smallest element that is > target
	for low < high-1 {
		mid := low + (high-low)/2
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid
		} else { // if target found, directly return the index
			return mid
		}
	}

	if nums[low] >= target {
		return low
	}

	if nums[high] >= target {
		return high
	}

	return high + 1
}

func searchInsertII(nums []int, target int) int {
	low, high := 0, len(nums)-1
	// range最后没有值不用post processing
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] > target {
			if mid == 0 || nums[mid-1] < target {
				return mid
			}
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			return target
		}
	}
	return -1
}

//func binarySearch() {
//	searchInsert([]int{1, 3, 5, 6}, 5)
//}
