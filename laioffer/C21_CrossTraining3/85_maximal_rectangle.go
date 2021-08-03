package main

// use histogram as sub problem
// TC: O(mn)
// SC: O(n)
func maximalRectangle85(matrix [][]byte) int {
	if matrix == nil || len(matrix) == 0 {
		return 0
	}
	input := make([]int, len(matrix[0]))
	max := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				input[j] += 1
			} else {
				input[j] = 0
			}
		}
		area := getMaxArea85(input)
		if area > max {
			max = area
		}
	}
	return max
}

func getMaxArea85(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	left[0] = -1
	right[n-1] = n

	for i := 1; i < n; i++ {
		prev := i - 1
		for prev >= 0 && heights[prev] >= heights[i] {
			// update prev
			// since idx = range (left[prev], prev] has heights[idx] >= heights[prev]
			// no need to check them, just skip
			prev = left[prev]
		}
		// now heights[prev] < heights[i] || prev = -1
		left[i] = prev
	}

	for i := n - 2; i >= 0; i-- {
		prev := i + 1
		for prev < n && heights[prev] >= heights[i] {
			// update prev
			// since idx = range [prev, right[prev]) has heights[idx] >= heights[prev]
			// no need to check them, just skip
			prev = right[prev]
		}
		// now heights[prev] < heights[i] || prev = n
		right[i] = prev
	}

	// find max area
	max := 0
	for i := 0; i < len(left); i++ {
		area := (right[i] - left[i] - 1) * heights[i]
		if max < area {
			max = area
		}
	}
	return max
}
