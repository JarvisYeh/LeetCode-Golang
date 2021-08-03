package main

func combine77(n int, k int) [][]int {
	res := &[][]int{}
	comb := &[]int{}
	dfs77(1, n, k, comb, res)
	return *res
}

func dfs77(start, end, remain int, comb *[]int, res *[][]int) {
	if remain == 0 {
		tmp := []int{}
		tmp = append(tmp, *comb...)
		*res = append(*res, tmp)
		return
	}

	for i := start; i <= end; i++ {
		*comb = append(*comb, i)
		dfs77(i+1, end, remain-1, comb, res)
		*comb = (*comb)[:len(*comb)-1]
	}
}
