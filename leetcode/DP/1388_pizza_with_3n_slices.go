package main

import "fmt"

func maxSizeSlices(slices []int) int {
	left := findMax(slices[:len(slices)-1])
	right := findMax(slices[1:])
	if left > right {
		return left
	} else {
		return right
	}
}

func findMax(slices []int) int {
	// sizes[i][j] is the max sizes of picking j slices among slices[0, i]
	sizes := make([][]int, len(slices))
	for i, _ := range sizes {
		sizes[i] = make([]int, (len(slices)+1)/3+1)
	}

	for i := 0; i < len(slices); i++ {
		for j := 1; j <= (len(slices)+1)/3 && j <= i+1; j++ {
			if i == 0 {
				sizes[i][j] = slices[i]
			} else if i == 1 { // if two elements, can only pick one for maximum
				if slices[0] > slices[1] {
					sizes[i][j] = slices[0]
				} else {
					sizes[i][j] = slices[1]
				}
			} else {
				// pick s[i]
				pick := sizes[i-2][j-1] + slices[i]
				// not pick s[i]
				notPick := sizes[i-1][j]
				if pick > notPick {
					sizes[i][j] = pick
				} else {
					sizes[i][j] = notPick
				}
			}
		}
	}
	return sizes[len(slices)-1][(len(slices)+1)/3]
}

func main() {
	fmt.Println(maxSizeSlices([]int{10, 9, 1, 10, 8, 5, 10, 2, 2}))
}
