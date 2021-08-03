package main

import (
	"math"
	"sort"
)

// method 1:
// binary search + two pointers
func findClosestElementsI(arr []int, k int, x int) []int {
	left, right := 0, len(arr)
	for left < right-1 {
		mid := left + ((right - left) >> 1)
		if arr[mid] == x {
			left = mid
			right = mid + 1
		} else if arr[mid] > x {
			right = mid
		} else {
			left = mid
		}
	}

	res := []int{}
	i, j := left, right
	for count := 0; count < k; count++ {
		if j >= len(arr) || (i >= 0 && math.Abs(float64(arr[i]-x)) <= math.Abs(float64(arr[j]-x))) {
			res = append(res, arr[i])
			i--
		} else {
			res = append(res, arr[j])
			j++
		}
	}
	sort.Ints(res)
	return res
}

// method 2:
// binary search + binary search in two sorted arrays
func findClosestElements(arr []int, k int, x int) []int {
	left, right := 0, len(arr)
	for left < right-1 {
		mid := left + ((right - left) >> 1)
		if arr[mid] == x {
			left = mid
			right = mid + 1
		} else if arr[mid] > x {
			right = mid
		} else {
			left = mid
		}
	}

	res := []int{}
	aLeft, bLeft := left, right
	for true {
		// base case 1: consume all numbers from left -> 0
		if aLeft < 0 {
			for i := 0; i < k; i++ {
				res = append(res, arr[bLeft+i])
			}
			break
		}
		// base case 2: consume all numbers from right -> len - 1
		if bLeft >= len(arr) {
			for i := 0; i < k; i++ {
				res = append(res, arr[aLeft-i])
			}
			break
		}
		// base case 3: only 1 number need to consider
		if k == 1 {
			if math.Abs(float64(arr[aLeft]-x)) <= math.Abs(float64(arr[bLeft]-x)) {
				res = append(res, arr[aLeft])
			} else {
				res = append(res, arr[bLeft])
			}
			break
		}

		aValue, bValue := math.MaxFloat64, math.MaxFloat64
		if aLeft-k/2+1 >= 0 {
			aValue = math.Abs(float64(arr[aLeft-k/2+1] - x))
		}

		if bLeft+k/2-1 < len(arr) {
			bValue = math.Abs(float64(arr[bLeft+k/2-1] - x))
		}
		if aValue < bValue {
			for i := 0; i < k/2; i++ {
				res = append(res, arr[aLeft-i])
			}
			aLeft -= k / 2
		} else {
			for i := 0; i < k/2; i++ {
				res = append(res, arr[bLeft+i])
			}
			bLeft += k / 2
		}
		k = k - k/2
	}
	sort.Ints(res)
	return res
}

func main() {
	findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3)
}
