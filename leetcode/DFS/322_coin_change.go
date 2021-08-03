package main

import (
	"fmt"
	"math"
)

func coinChange322(coins []int, amount int) int {
	min := new(int)
	*min = math.MaxInt64
	dfs322(0, 0, min, coins, amount)
	if *min == math.MaxInt64 {
		return -1
	}
	return *min
}

func dfs322(coinIndex, coinAmount int, min *int, coins []int, remain int) {
	// base case
	if coinIndex == len(coins)-1 {
		if remain%coins[coinIndex] == 0 {
			if coinAmount+remain/coins[coinIndex] < *min {
				*min = coinAmount
			}
		}
		return
	}

	for i := 0; i <= remain/coins[coinIndex]; i++ {
		dfs322(coinIndex+1, coinAmount+i, min, coins, remain-i*coins[coinIndex])
	}
}

func main() {
	fmt.Println(coinChange322([]int{411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422}, 9864))
}
