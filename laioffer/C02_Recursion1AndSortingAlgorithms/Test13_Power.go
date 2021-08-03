package main

/**
 * Assume no overflows and b >= 0
 */
func pow13I(a int, b int) int {
	if b == 0 {
		return 1
	}

	tmp := pow13I(a, b/2)

	if b%2 == 0 {
		return tmp * tmp
	} else {
		return tmp * tmp * b
	}
}

// TC: O((log n)^2)
//	worst case, e.g. 1023 = 1 + 2 + 4 + ... + 512
// 	O(log n) + O((log n) - 1) + O((log n) - 2) + O((log n) - 3) + ... + O(1) = O((log n)^2)
// SC: O(1)
func pow13II(a int, b int) int {
	if b == 0 {
		return 1
	}

	result, remain := 1, b
	for remain > 0 {
		pow, count := a, 1
		for (count << 1) < remain {
			pow *= pow
			count <<= 1
		}
		result *= pow
		remain -= count
	}
	return result
}
