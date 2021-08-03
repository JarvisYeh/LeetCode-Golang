package main

import "fmt"

// DFS 1:
// each level in recursion tree represents cut/not cut
// there is total n - 1 gap to be cut
// n - 1 levels
// each level two nodes represents cut/not cut
// TC: O(2^n)
// SC: O(n) in stack
func integerBreak343I(n int) int {
	return findMaxProd343I(1, 1, n, false)
}

// reverse is the current length accumulated so far, including the 1 before this position gap
// __|__ position = 2, reverse = 2
// this function returns max product of length n
func findMaxProd343I(position, reverse, length int, checkCut bool) int {
	// corner case, can not divide 1
	if length == 1 {
		return -1
	}
	// base case: last gap
	if position == length-1 {
		if checkCut {
			return reverse + 1
		} else {
			return reverse * 1
		}
	}

	// if cut, the reverse amount is settled now, and reverse back to 1
	cut := reverse * findMaxProd343I(position+1, 1, length, true)
	// if not cut, call recursive function
	notCut := findMaxProd343I(position+1, reverse+1, length, checkCut)
	if cut > notCut {
		return cut
	} else {
		return notCut
	}
}

// DFS 2:
// each level has n - 1 nodes, represents cut in each gap
// TC: O(2^n)
// SC: O(n) in stack
func integerBreak343II(n int) int {
	// base case: 1 meter rope can not be cut, return 1
	if n <= 1 {
		return 1
	}

	prod := 0
	// left : cut i-th position, left part is settled i meters rope
	// right : max(integerBreak(n-i), n-i)
	//      e.g. max right produc if cut, or the whole right part not cut
	for i := 1; i < n; i++ {
		currProd := i * (n - i)
		rightProd := integerBreak343II(n - i)
		if rightProd > n-i {
			currProd = currProd / (n - i) * rightProd
		}
		if currProd > prod {
			prod = currProd
		}
	}
	return prod
}

// DP 1
// both left and right segment
// need to be looked at dict
// TC: O(n^2)
// SC: O(n)
func integerBreak343III(n int) int {
	// prods[i] stores the max product for rope with length n (at least one cut)
	prods := make([]int, n+1)
	prods[1] = 1 // actually prod[1] is invalid, but max(prods[1], 1) will give 1, so doesn't matter

	for i := 2; i <= n; i++ {
		maxProd := 0
		for j := 1; j < i; j++ {
			// left = max(left, prods[left])
			left := j
			if prods[left] > left {
				left = prods[left]
			}
			// right = max(right, prods[right])
			right := i - j
			if prods[right] > right {
				right = prods[right]
			}

			// maxProd = max(maxProd, left*right)
			if maxProd < left*right {
				maxProd = left * right
			}
		}
		prods[i] = maxProd
	}
	return prods[n]
}

// DP 2
// just check the left segment from dict
func integerBreak343IV(n int) int {
	// prods[i] stores the max product for rope with length n (at least one cut)
	prods := make([]int, n+1)
	prods[1] = 1 // actually prod[1] is invalid

	for i := 2; i <= n; i++ {
		maxProd := 0
		for j := 1; j < i; j++ {
			// left = max(left, prods[left])
			left := j
			if prods[left] > left {
				left = prods[left]
			}
			// right = i - j
			// cause when j increases, there will be some time that left = i - j and right = j
			right := i - j

			// maxProd = max(maxProd, left*right)
			if maxProd < left*right {
				maxProd = left * right
			}
		}
		prods[i] = maxProd
	}
	return prods[n]
}

func main() {
	fmt.Println(integerBreak343IV(30))
}
