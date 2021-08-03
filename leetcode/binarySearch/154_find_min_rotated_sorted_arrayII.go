package main

func findMin154(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		// if [low, high] is already sorted, directly return the first element
		if nums[low] < nums[high] {
			return nums[low]
		}

		// else there is rotation exist
		mid := low + ((high - low) >> 1)
		if nums[mid] > nums[high] || nums[mid] > nums[low] { // [low mid | high]
			low = mid + 1
		} else if nums[mid] < nums[high] || nums[mid] < nums[low] { // [low | mid high]
			high = mid
		} else { //nums[mid] == nums[high] == nums[low]
			high-- // shrink boundary in case of infinite loop
		}
	}
	// after loop, lwo = high
	return nums[low]
}

func main() {
	findMin154([]int{1, 3, 5})
}
