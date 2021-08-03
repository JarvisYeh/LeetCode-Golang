package main

import "fmt"

func rotate48(matrix [][]int) {
	helper48(0, len(matrix), matrix)
}

func helper48(offset, size int, mat [][]int) {
	if size <= 1 {
		return
	}

	for i := 0; i < size-1; i++ {
		// tmp = top side
		tmp := mat[offset][offset+i]
		// left side to top side
		mat[offset][offset+i] = mat[offset+size-1-i][offset]
		// bottom side to left side
		mat[offset+size-1-i][offset] = mat[offset+size-1][offset+size-1-i]
		// right side to bottom side
		mat[offset+size-1][offset+size-1-i] = mat[offset+i][offset+size-1]
		// top side to right side
		mat[offset+i][offset+size-1] = tmp
	}

	helper48(offset+1, size-2, mat)
}

func main() {
	tmp := [][]int{{1, 2}, {3, 4}}
	rotate48(tmp)
	fmt.Println(tmp)
}
