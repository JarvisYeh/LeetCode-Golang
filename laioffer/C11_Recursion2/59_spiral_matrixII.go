package main

import "fmt"

func generateMatrix59(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	traverse59(res, 0, n, 1)
	return res
}

func traverse59(res [][]int, offset, size, counter int) {
	// base case: 1*1 square or 0*0 square
	if size <= 1 {
		if size == 1 {
			res[offset][offset] = counter
		}
		return
	}

	// top start point : [offset, offset]
	for i, j := offset, offset; j < offset+size-1; j++ {
		res[i][j] = counter
		counter++
	}

	// right start point : [offset, offset + size - 1]
	for i, j := offset, offset+size-1; i < offset+size-1; i++ {
		res[i][j] = counter
		counter++
	}

	// bottom start point : [offset + size - 1, offset + size - 1]
	for i, j := offset+size-1, offset+size-1; j > offset; j-- {
		res[i][j] = counter
		counter++
	}

	// left start point : [offset + size - 1, offset]
	for i, j := offset+size-1, offset; i > offset; i-- {
		res[i][j] = counter
		counter++
	}

	traverse59(res, offset+1, size-2, counter)
}

func generateMatrix59Iter(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	offset, counter := 0, 1
	size := n
	for ; size > 1; size -= 2 {
		// top start point : [offset, offset]
		for i, j := offset, offset; j < offset+size-1; j++ {
			res[i][j] = counter
			counter++
		}

		// right start point : [offset, offset + size - 1]
		for i, j := offset, offset+size-1; i < offset+size-1; i++ {
			res[i][j] = counter
			counter++
		}

		// bottom start point : [offset + size - 1, offset + size - 1]
		for i, j := offset+size-1, offset+size-1; j > offset; j-- {
			res[i][j] = counter
			counter++
		}

		// left start point : [offset + size - 1, offset]
		for i, j := offset+size-1, offset; i > offset; i-- {
			res[i][j] = counter
			counter++
		}
		offset += 1
	}

	if size == 1 {
		res[offset][offset] = counter
	}
	return res
}

func main() {
	res := generateMatrix59Iter(6)
	for _, val := range res {
		fmt.Println(val)
	}

}
