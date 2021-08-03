package main

func searchMatrix(matrix [][]int, target int) bool {
	row, col := len(matrix), len(matrix[0])
	low, high := 0, row*col-1
	var i, j int
	for low <= high {
		mid := low + ((high - low) >> 1)
		i = mid / col
		j = mid % col
		if matrix[i][j] < target {
			low = mid + 1
		} else if matrix[i][j] > target {
			high = mid - 1
		} else {
			return true
		}
	}
	return false
}
