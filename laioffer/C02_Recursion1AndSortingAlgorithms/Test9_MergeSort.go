package main

// TC: O(nlogn)
// SC: O(n)
func mergeSort(nums []int, left, right int) []int {
	if left == right {
		return []int{nums[left]}
	}

	mid := left + ((right - left) >> 1)
	leftArr := mergeSort(nums, left, mid)
	rightArr := mergeSort(nums, mid+1, right)
	// merge
	return merge(leftArr, rightArr)
}

func merge(left, right []int) []int {
	res := make([]int, len(left)+len(right))
	leftIndex, rightIndex, resIndex := 0, 0, 0
	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			res[resIndex] = left[leftIndex]
			leftIndex++
		} else {
			res[resIndex] = right[rightIndex]
			rightIndex++
		}
		resIndex++
	}

	for leftIndex < len(left) {
		res[resIndex] = left[leftIndex]
		leftIndex++
		resIndex++
	}

	for rightIndex < len(right) {
		res[resIndex] = right[rightIndex]
		rightIndex++
		resIndex++
	}
	return res
}

// mergeSort with one helper array
// every time it merges [left, mid] and [mid + 1, right]
// it use one helper array as reference and put the right number into nums array
func mergeSortWithHelperArray(nums []int) []int {
	helper := make([]int, len(nums))
	mergeSortWithHelper(nums, helper, 0, len(nums)-1)
	return helper
}

func mergeSortWithHelper(nums, helper []int, left, right int) {
	if left == right {
		return
	}
	mid := left + ((right - left) >> 1)
	mergeSortWithHelper(nums, helper, left, mid)
	mergeSortWithHelper(nums, helper, mid+1, right)
	mergeWithHelper(nums, helper, left, mid, right)
}

func mergeWithHelper(nums, helper []int, left, mid, right int) {
	// copy [left, right] from nums to helper
	for i := left; i <= right; i++ {
		helper[i] = nums[i]
	}

	leftIndex, rightIndex, resIndex := left, mid+1, left
	for leftIndex <= mid && rightIndex <= right {
		if helper[leftIndex] < helper[rightIndex] {
			nums[resIndex] = helper[leftIndex]
			leftIndex++
		} else {
			nums[resIndex] = helper[rightIndex]
			rightIndex++
		}
		resIndex++
	}

	for leftIndex <= mid {
		nums[resIndex] = helper[leftIndex]
		resIndex++
		leftIndex++
	}
}

func main() {
	mergeSortWithHelperArray([]int{5, 1, 1, 2, 0, 0})
}
