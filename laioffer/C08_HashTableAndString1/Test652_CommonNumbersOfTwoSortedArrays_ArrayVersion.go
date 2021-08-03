package main

// common elements in two sorted arrays
// no duplication

// use hashmap
// TC: O(m + n)
// SC: O(min(m, n))
func common652I(nums1, nums2 []int) []int {
	res := []int{}
	m := make(map[int]bool)

	var smallArr, largeArr []int
	if len(nums1) < len(nums2) {
		smallArr = nums1
		largeArr = nums2
	} else {
		smallArr = nums2
		largeArr = nums1
	}

	// store info of smaller arr into map
	for _, num := range smallArr {
		if !m[num] {
			m[num] = true
		}
	}

	// check each element in large arr
	for _, num := range largeArr {
		if m[num] {
			res = append(res, num)
		}
	}
	return res
}

// use two pointer
// TC: O(m + n)
// SC: O(1)
func common652II(nums1, nums2 []int) []int {
	res := []int{}
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			res = append(res, nums1[i])
			i++
			j++
		}
	}
	return res
}

// use binary search
// TC: O(min(mlogn, nlogm))
// if n >>>> m
// log n can be shorten well, mlogn < m + n, better than two pointer method
// SC: O(1)
func common652III(nums1, nums2 []int) []int {
	res := []int{}

	var smallArr, largeArr []int
	if len(nums1) < len(nums2) {
		smallArr = nums1
		largeArr = nums2
	} else {
		smallArr = nums2
		largeArr = nums1
	}

	// search every ele in smaller array among the large array
	for _, num := range smallArr {
		if binarySearch652(num, largeArr) {
			res = append(res, num)
		}
	}
	return res
}

func binarySearch652(target int, arr []int) bool {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] > target {
			r = mid - 1
		} else if arr[mid] < target {
			l = mid + 1
		} else {
			return true
		}
	}
	return false
}
