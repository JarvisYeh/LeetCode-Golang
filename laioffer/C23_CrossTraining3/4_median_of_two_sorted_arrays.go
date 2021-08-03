package main

// TC: O(log k)
// SC: O(log k)
func findMedianSortedArrays4(nums1 []int, nums2 []int) float64 {
	lenSum := len(nums1) + len(nums2)
	if lenSum%2 == 0 {
		left := findK4(nums1, nums2, 0, 0, lenSum/2)
		right := findK4(nums1, nums2, 0, 0, lenSum/2+1)
		return float64(left+right) / 2
	} else {
		return float64(findK4(nums1, nums2, 0, 0, lenSum/2+1))
	}
}

// k is the number of elements <= target INCLUDE target
func findK4(nums1, nums2 []int, l1, l2, k int) int {
	// base cases
	if l1 >= len(nums1) {
		return nums2[l2+k-1]
	}
	if l2 >= len(nums2) {
		return nums1[l1+k-1]
	}
	if k == 1 {
		if nums1[l1] < nums2[l2] {
			return nums1[l1]
		} else {
			return nums2[l2]
		}
	}

	// first set mid1 and mid2 to max int
	mid1, mid2 := int(^uint(0)>>1), int(^uint(0)>>1)
	// if mid is in the range of array, assign its correct value
	if l1+k/2-1 < len(nums1) {
		mid1 = nums1[l1+k/2-1]
	}
	if l2+k/2-1 < len(nums2) {
		mid2 = nums2[l2+k/2-1]
	}

	// compare two mid
	if mid1 < mid2 {
		return findK4(nums1, nums2, l1+k/2, l2, k-k/2)
	} else {
		return findK4(nums1, nums2, l1, l2+k/2, k-k/2)
	}

}
