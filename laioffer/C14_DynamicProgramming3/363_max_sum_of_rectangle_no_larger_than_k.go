package main

import "math"

// 2D DP
// TC: O(n^4)
// SC: O(n^2)
func maxSumSubmatrix(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	// prefixSum[i][j] represents the sum of square
	// with right down corner as (i, j)
	prefixSum := make([][]int, m)
	for i := 0; i < m; i++ {
		prefixSum[i] = make([]int, n)
	}

	// construct prefixSum[][]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				prefixSum[i][j] = matrix[i][j]
			} else if i == 0 {
				prefixSum[i][j] = prefixSum[i][j-1] + matrix[i][j]
			} else if j == 0 {
				prefixSum[i][j] = prefixSum[i-1][j] + matrix[i][j]
			} else {
				prefixSum[i][j] = prefixSum[i-1][j] + prefixSum[i][j-1] - prefixSum[i-1][j-1] + matrix[i][j]
			}
		}
	}

	max := math.MinInt64
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for x := i; x < m; x++ {
				for y := j; y < n; y++ {
					currMax := 0
					if i == 0 && j == 0 {
						currMax = prefixSum[x][y]
					} else if i == 0 {
						currMax = prefixSum[x][y] - prefixSum[x][j-1]
					} else if j == 0 {
						currMax = prefixSum[x][y] - prefixSum[i-1][y]
					} else {
						currMax = prefixSum[x][y] - prefixSum[x][j-1] - prefixSum[i-1][y] + prefixSum[i-1][j-1]
					}
					if currMax <= k && currMax > max {
						max = currMax
					}
				}
			}
		}
	}
	if max == math.MinInt64 {
		return -1
	} else {
		return max
	}
}
