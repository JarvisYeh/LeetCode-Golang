package main

// TC : O(n)
// SC: O(1)
// start with i = 0, j = len(height) - 1
// height = min(i, j)
// to increase the volumn of water, we either increase height or increase width
// but now width is at maximum, the only possible way to go is scarifies some width with much larger height
// since height = min(h[i], h[j])
// only move the smaller from h[i], h[j] can possibly achieve higher height
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	max := 0
	for i < j {
		// h = min(h[i], h[j])
		h := height[i]
		if h > height[j] {
			h = height[j]
		}

		if max < h*(j-i) {
			max = h * (j - i)
		}

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}
