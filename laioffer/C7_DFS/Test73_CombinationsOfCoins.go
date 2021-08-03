package main

import "fmt"

func combinations73(target int, coins []int) [][]int {
	res := &[][]int{}
	comb := &[]int{}
	findComb73(0, target, coins, comb, res)
	return *res
}

func findComb73(index, rem int, coins []int, comb *[]int, res *[][]int) {
	if index == len(coins)-1 {
		if rem%coins[index] == 0 {
			*comb = append(*comb, rem/coins[index])
			tmp := []int{}
			tmp = append(tmp, *comb...)
			*res = append(*res, tmp)
			*comb = (*comb)[:len(*comb)-1]
		}
		return
	}

	for i := 0; i <= rem/coins[index]; i++ {
		*comb = append(*comb, i)
		findComb73(index+1, rem-coins[index]*i, coins, comb, res)
		*comb = (*comb)[:len(*comb)-1]
	}
}

func main() {
	fmt.Println(combinations73(10, []int{34, 31, 29, 16, 2}))
}
