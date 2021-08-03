package main

// TC : O(n^3)
// SC: O(n^2)
func orderOfLargestPlusSignI(n int, mines [][]int) int {
	// create the matrix: O(n^2)
	mat := make([][]int, n)
	for i, _ := range mat {
		mat[i] = make([]int, n)
		for j, _ := range mat[i] {
			mat[i][j] = 1
		}
	}

	// add 0s, O(n^2)
	for _, val := range mines {
		mat[val[0]][val[1]] = 0
	}

	// for each point check 4 directions
	// O(n^3)
	max := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] != 1 {
				continue
			}
			up := 0
			// up
			for i-up >= 0 && mat[i-up][j] == 1 {
				up++
			}
			right := 0
			for j+right < n && mat[i][j+right] == 1 {
				right++
			}
			down := 0
			for i+down < n && mat[i+down][j] == 1 {
				down++
			}
			left := 0
			for j-left >= 0 && mat[i][j-left] == 1 {
				left++
			}
			min := up
			if right < min {
				min = right
			}
			if left < min {
				min = left
			}
			if down < min {
				min = down
			}
			if min > max {
				max = min
			}
		}
	}
	return max
}

// DP
// get the consecutive 1s from [i][j] to left, right, up, down
// mat[i][j] = min(left, right, up, down)
// max = max of all min(left, right, up, down)
// TC: O(5*n^2)
// SC: O(n^2)
func orderOfLargestPlusSign(n int, mines [][]int) int {
	mat := make([][]int, n)
	for i, _ := range mat {
		mat[i] = make([]int, n)
		for j, _ := range mat[i] {
			mat[i][j] = 1
		}
	}

	for _, val := range mines {
		mat[val[0]][val[1]] = 0
	}

	// mat[i][j] = min(left, right)
	for i := 0; i < n; i++ {
		countLeft := 0
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				countLeft = 0
			} else {
				countLeft++
				mat[i][j] = countLeft
			}
		}
		countRight := 0
		for j := n - 1; j >= 0; j-- {
			if mat[i][j] == 0 {
				countRight = 0
			} else {
				countRight++
				if countRight < mat[i][j] {
					mat[i][j] = countRight
				}
			}
		}
	}

	// mat[i][j] = min(mat[i][j], up, down)
	for j := 0; j < n; j++ {
		countUp := 0
		for i := 0; i < n; i++ {
			if mat[i][j] == 0 {
				countUp = 0
			} else {
				countUp++
				if countUp < mat[i][j] {
					mat[i][j] = countUp
				}
			}
		}
		countDown := 0
		for i := n - 1; i >= 0; i-- {
			if mat[i][j] == 0 {
				countDown = 0
			} else {
				countDown++
				if countDown < mat[i][j] {
					mat[i][j] = countDown
				}
			}
		}
	}

	max := 0
	// max = max(min(left, right, up, down))
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] > max {
				max = mat[i][j]
			}
		}
	}
	return max
}
