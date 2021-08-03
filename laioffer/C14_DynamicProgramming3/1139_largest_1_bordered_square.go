package main

func largest1BorderedSquare1139(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	left := make([][]int, m)
	up := make([][]int, m)
	for i := 0; i < m; i++ {
		left[i] = make([]int, n)
		up[i] = make([]int, n)
	}

	// constuct up and left matrix
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && j-1 >= 0 {
				left[i][j] = left[i][j-1] + 1
			} else {
				left[i][j] = grid[i][j]
			}

			if grid[i][j] == 1 && i-1 >= 0 {
				up[i][j] = up[i-1][j] + 1
			} else {
				up[i][j] = grid[i][j]
			}
		}
	}

	gMaxSize := 0
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			maxSize := min1139(left[i][j], up[i][j])
			if maxSize < gMaxSize {
				continue
			}
			for size := maxSize; size >= 1; size-- {
				// check up[leftDown corner] and left[rightUp corner]
				if up[i][j-size+1] >= size && left[i-size+1][j] >= size {
					gMaxSize = max1139(size, gMaxSize)
					break
				}
			}
		}
	}
	return gMaxSize * gMaxSize
}

func max1139(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func min1139(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}
