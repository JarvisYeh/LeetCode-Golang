package main

import "fmt"

func validParentheses642(l, m, n int) []string {
	stack := &[]byte{}
	left := []int{0, 0, 0}
	right := []int{0, 0, 0}
	res := &[]string{}
	generateParentheses642(left, right, stack, l, m, n, []byte{}, res)
	return *res
}

func generateParentheses642(left, right []int, stack *[]byte, l, m, n int, curr []byte, res *[]string) {
	if left[0]+right[0] == 2*l && left[1]+right[1] == 2*m && left[2]+right[2] == 2*n {
		*res = append(*res, string(curr))
		return
	}

	// () can be in {} and <>
	if left[0] < l && (len(*stack) == 0 || (*stack)[len(*stack)-1] == '{' || (*stack)[len(*stack)-1] == '<') {
		left[0]++
		*stack = append(*stack, '(')
		generateParentheses642(left, right, stack, l, m, n, append(curr, '('), res)
		*stack = (*stack)[:len(*stack)-1]
		left[0]--
	}

	// <> can only be in {}
	if left[1] < m && (len(*stack) == 0 || (*stack)[len(*stack)-1] == '{') {
		left[1]++
		*stack = append(*stack, '<')
		generateParentheses642(left, right, stack, l, m, n, append(curr, '<'), res)
		*stack = (*stack)[:len(*stack)-1]
		left[1]--
	}

	// {} can only be in the outside
	if left[2] < n && len(*stack) == 0 {
		left[2]++
		*stack = append(*stack, '{')
		generateParentheses642(left, right, stack, l, m, n, append(curr, '{'), res)
		*stack = (*stack)[:len(*stack)-1]
		left[2]--
	}

	if len(*stack) > 0 && (*stack)[len(*stack)-1] == '(' {
		right[0]++
		*stack = (*stack)[:len(*stack)-1]
		generateParentheses642(left, right, stack, l, m, n, append(curr, ')'), res)
		*stack = append(*stack, '(')
		right[0]--
	}

	if len(*stack) > 0 && (*stack)[len(*stack)-1] == '<' {
		right[1]++
		*stack = (*stack)[:len(*stack)-1]
		generateParentheses642(left, right, stack, l, m, n, append(curr, '>'), res)
		*stack = append(*stack, '<')
		right[1]--
	}

	if len(*stack) > 0 && (*stack)[len(*stack)-1] == '{' {
		right[2]++
		*stack = (*stack)[:len(*stack)-1]
		generateParentheses642(left, right, stack, l, m, n, append(curr, '}'), res)
		*stack = append(*stack, '{')
		right[2]--
	}
}

func main() {
	for _, s := range validParentheses642(1, 1, 0) {
		fmt.Println(s)
	}
}
