package main

import "sort"

// two unsorted array, return all intersection elements
// including duplication

func intersect350I(nums1 []int, nums2 []int) []int {
	res := []int{}
	m := make(map[int]int)
	// store the frequency of nums1
	// O(n)
	for _, n := range nums1 {
		if _, ok := m[n]; ok {
			m[n]++
		} else {
			m[n] = 1
		}
	}

	// check nums2, O(m)
	for _, n := range nums2 {
		if count, ok := m[n]; ok {
			if count > 0 {
				res = append(res, n)
				m[n]--
			}
		}
	}

	return res
}

func intersect350II(nums1 []int, nums2 []int) []int {
	res := []int{}
	// O(nlogn) + O(mlogm)
	sort.Ints(nums1)
	sort.Ints(nums2)
	i, j := 0, 0
	// O(min(n, m))
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i++
			j++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}
	}
	return res
}
