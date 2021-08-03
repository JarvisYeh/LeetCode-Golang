package main

import (
	"fmt"
	"math"
)

func MIDTERM_nQueens(n int) [][]int {
	res := make([][]int, 0)
	curr := make([]int, 0)
	MIDTERM_findNQueenComb(curr, &res, n)
	return res
}

func MIDTERM_findNQueenComb(curr []int, res *[][]int, n int) {
	if len(curr) == n {
		tmp := make([]int, n)
		copy(tmp, curr)
		*res = append(*res, tmp)
		return
	}

	for c := 0; c < n; c++ {
		if MIDTERM_checkValid(c, curr) {
			curr = append(curr, c)
			MIDTERM_findNQueenComb(curr, res, n)
			curr = curr[:len(curr)-1]
		}
	}
}

func MIDTERM_checkValid(col int, curr []int) bool {
	row := len(curr)
	for r := 0; r < len(curr); r++ {
		c := curr[r]
		if col == c || row-r == int(math.Abs(float64(col)-float64(c))) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(MIDTERM_nQueens(5))
}
