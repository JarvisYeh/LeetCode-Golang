package main

// two unsorted array, return "unique" intersection elements
// including duplication

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]bool)
	res := []int{}
	for _, n := range nums1 {
		if !m[n] {
			m[n] = true
		}
	}

	for _, n := range nums2 {
		if m[n] {
			res = append(res, n)
			delete(m, n)
		}
	}
	return res
}
