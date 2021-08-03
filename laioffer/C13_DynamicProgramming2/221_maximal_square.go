package main

// return the max "area" of all '1's square in matrix
func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])

	// len[i][j] represents the max length of all 1s squares
	// whose right bottom corner is (i, j)
	lens := make([][]int, m)
	for i, _ := range lens {
		lens[i] = make([]int, n)
	}

	max := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// base case, top and left side
			// side length equals to matrix[i][j]
			if i == 0 || j == 0 {
				lens[i][j] = int(matrix[i][j] - '0')
			} else if matrix[i][j] == '1' {
				// if it's not the left and top side of matrix
				// and mat[i][j] contains a '1'
				// check the size of all 1s square of top, left, and topLeft
				top := lens[i-1][j]
				left := lens[i][j-1]
				topLeft := lens[i-1][j-1]
				min := top
				if left < min {
					min = left
				}
				if topLeft < min {
					min = topLeft
				}
				lens[i][j] = min + 1
			}

			if lens[i][j] > max {
				max = lens[i][j]
			}
		}
	}
	return max * max
}
