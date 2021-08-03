package main

import "fmt"

func generateParenthesis22(n int) []string {
	res := &[]string{}
	comb := &[]rune{}
	dfs22(0, 0, n, comb, res)
	return *res
}

func dfs22(left, right, n int, comb *[]rune, res *[]string) {
	if left+right == 2*n {
		*res = append(*res, string(*comb))
		return
	}

	if left < n {
		*comb = append(*comb, '(')
		dfs22(left+1, right, n, comb, res)
		*comb = (*comb)[:len(*comb)-1]
	}

	if right < left {
		*comb = append(*comb, ')')
		dfs22(left, right+1, n, comb, res)
		*comb = (*comb)[:len(*comb)-1]
	}
}

func main() {
	fmt.Println(generateParenthesis22(2))
}
