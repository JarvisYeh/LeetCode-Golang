package main

import "math"

func totalNQueens52(n int) int {
	res := new(int)
	*res = 0
	comb := []int{}
	dfs52(comb, res, n)
	return *res
}

func dfs52(comb []int, res *int, n int) {
	if len(comb) == n {
		*res += 1
		return
	}

	for i := 0; i < n; i++ {
		if checkPass52(comb, i) {
			comb = append(comb, i)
			dfs52(comb, res, n)
			comb = comb[:len(comb)-1]
		}
	}
}

func checkPass52(comb []int, col int) bool {
	row := len(comb)
	for r := 0; r < len(comb); r++ {
		c := comb[r]
		if col == c || math.Abs(float64(row)-float64(r)) == math.Abs(float64(col)-float64(c)) {
			return false
		}
	}
	return true
}

func main() {
	totalNQueens52(4)
}
