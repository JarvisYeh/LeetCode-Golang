package main

import (
	"fmt"
	"strconv"
)

func MIDTERM_printNPairOfIf(n int) {
	combs := make([][]rune, 0)
	curr := []rune{}
	MIDTERM_getComb(0, 0, n, curr, &combs)
	MIDTERM_printComb(combs)
}

func MIDTERM_getComb(left, right, n int, curr []rune, combs *[][]rune) {
	if left == n && right == n {
		tmp := make([]rune, 2*n)
		copy(tmp, curr)
		*combs = append(*combs, tmp)
		return
	}

	if left < n {
		curr = append(curr, '{')
		MIDTERM_getComb(left+1, right, n, curr, combs)
		curr = curr[:len(curr)-1]
	}

	if right < left {
		curr = append(curr, '}')
		MIDTERM_getComb(left, right+1, n, curr, combs)
		curr = curr[:len(curr)-1]
	}
}

func MIDTERM_printComb(combs [][]rune) {
	stack := []rune{}
	for i, comb := range combs {
		fmt.Println(strconv.Itoa(i) + "-th combination")
		// spaces is the amount for next level if the start a new '{'
		spaces := 0
		for _, val := range comb {
			if val == '{' {
				stack = append(stack, '{')
				for i := 0; i < spaces; i++ {
					fmt.Print(" ")
				}
				fmt.Println("if {")
				spaces += 2
			} else {
				stack = stack[:len(stack)-1]
				spaces -= 2
				for i := 0; i < spaces; i++ {
					fmt.Print(" ")
				}
				fmt.Println("}")
			}
		}
		fmt.Println("")
	}
}

func main() {
	MIDTERM_printNPairOfIf(4)
}
