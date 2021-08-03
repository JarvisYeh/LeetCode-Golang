package main

import "fmt"

func validParentheses179(l, m, n int) []string {
	stack := &[]byte{}
	left := []int{0, 0, 0}
	right := []int{0, 0, 0}
	res := &[]string{}
	generateParentheses179(left, right, stack, l, m, n, []byte{}, res)
	return *res
}

func generateParentheses179(left, right []int, stack *[]byte, l, m, n int, curr []byte, res *[]string) {
	if left[0]+right[0] == 2*l && left[1]+right[1] == 2*m && left[2]+right[2] == 2*n {
		*res = append(*res, string(curr))
		return
	}

	if left[0] < l {
		left[0]++
		*stack = append(*stack, '(')
		generateParentheses179(left, right, stack, l, m, n, append(curr, '('), res)
		*stack = (*stack)[:len(*stack)-1]
		left[0]--
	}

	if left[1] < m {
		left[1]++
		*stack = append(*stack, '<')
		generateParentheses179(left, right, stack, l, m, n, append(curr, '<'), res)
		*stack = (*stack)[:len(*stack)-1]
		left[1]--
	}

	if left[2] < n {
		left[2]++
		*stack = append(*stack, '{')
		generateParentheses179(left, right, stack, l, m, n, append(curr, '{'), res)
		*stack = (*stack)[:len(*stack)-1]
		left[2]--
	}

	if len(*stack) > 0 && (*stack)[len(*stack)-1] == '(' {
		right[0]++
		*stack = (*stack)[:len(*stack)-1]
		generateParentheses179(left, right, stack, l, m, n, append(curr, ')'), res)
		*stack = append(*stack, '(')
		right[0]--
	}

	if len(*stack) > 0 && (*stack)[len(*stack)-1] == '<' {
		right[1]++
		*stack = (*stack)[:len(*stack)-1]
		generateParentheses179(left, right, stack, l, m, n, append(curr, '>'), res)
		*stack = append(*stack, '<')
		right[1]--
	}

	if len(*stack) > 0 && (*stack)[len(*stack)-1] == '{' {
		right[2]++
		*stack = (*stack)[:len(*stack)-1]
		generateParentheses179(left, right, stack, l, m, n, append(curr, '}'), res)
		*stack = append(*stack, '{')
		right[2]--
	}
}

func main() {
	for _, s := range validParentheses179(1, 1, 1) {
		fmt.Println(s)
	}
}
