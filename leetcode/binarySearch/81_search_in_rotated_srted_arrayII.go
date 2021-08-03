package main

func search81(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2

		// found and return directly
		if nums[mid] == target {
			return mid
		}

		// two segment [low | high], if rotation > 0, nums[low] = nums[high] + c
		if nums[mid] > nums[low] { // [low mid | high]
			if target >= nums[low] && target < nums[mid] { // [low t mid | high]
				high = mid - 1
			} else { // [low mid | t high] or [low mid t | high]
				low = mid + 1
			}
		} else if nums[mid] < nums[high] { // [low | mid high]
			if target > nums[mid] && target <= nums[high] { // [low | mid t high]
				low = mid + 1
			} else { // [low | t mid high] or [low t | mid high]
				high = mid - 1
			}
		} else { // nums[high] <= nums[mid] <= nums[low]
			// since nums[low] is the next bigger one than nums[high] in array, there is no element between low and high
			// it can only be nums[mid] == nums[low] != target || nums[mid] == nums[high] != target
			// which ever the case, we have to shrink the boundary by 1 to avoid infinite loop
			if nums[mid] == nums[low] {
				low++
			} else {
				high--
			}
		}
	}
	return -1
}
