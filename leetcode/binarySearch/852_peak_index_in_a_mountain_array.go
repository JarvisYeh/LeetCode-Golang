package main

func peakIndexInMountainArray852(arr []int) int {
	l, r := 0, len(arr)-1
	for l < r-1 {
		mid := l + (r-l)/2
		if arr[mid-1] > arr[mid] {
			r = mid
		} else if arr[mid+1] > arr[mid] {
			l = mid
		} else if arr[mid+1] < arr[mid] && arr[mid-1] < arr[mid] {
			return mid
		}
	}

	if arr[l] > arr[r] {
		return l
	} else {
		return r
	}
}
