package main

import (
	"math"
)

// TC: O(log (dividen/divisor))
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	sign := -1
	if dividend > 0 && divisor > 0 || dividend < 0 && divisor < 0 {
		sign = 1
	}

	// set dividend and divisor to negative number
	// so that when dividend is min_int, it won't cause error
	if divisor > 0 {
		divisor = -divisor
	}

	if dividend > 0 {
		dividend = -dividend
	}

	var remain, result int
	remain = dividend
	result = 0
	for remain <= divisor {
		sum := divisor
		count := 1
		for sum+sum < 0 && sum+sum > remain {
			sum += sum
			count += count
		}
		remain -= sum
		result += count
	}

	if sign > 0 {
		return result
	} else {
		return -result
	}

}
