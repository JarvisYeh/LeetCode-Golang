package main

func mySqrt(x int) int {
	// set 0 as corner case
	if x == 0 {
		return 0
	}

	// binary search starts from 1 to avoid divided by 0 issue
	// use division rather than multiplication to avoid overflows
	low, high := 1, x
	for low <= high {
		mid := low + (high-low)/2
		if mid > x/mid {
			high = mid - 1
		} else if mid < x/mid {
			if (mid + 1) > x/(mid+1) {
				return mid
			}
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

//func binarySearch() {
//	fmt.Println(mySqrt(math.MaxInt64))
//}
