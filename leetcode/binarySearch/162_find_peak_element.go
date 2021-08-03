package main

func findPeakElement(nums []int) int {
	low, high := 0, len(nums)-1

	for low < high {
		mid := low + ((high - low) >> 1)
		if mid+1 < len(nums) && nums[mid] > nums[mid+1] {
			high = mid
		} else if mid+1 < len(nums) && nums[mid] < nums[mid+1] {
			low = mid + 1
		}
	}
	return nums[low]
}

func main() {
	findPeakElement([]int{1, 2, 3, 1})
}
