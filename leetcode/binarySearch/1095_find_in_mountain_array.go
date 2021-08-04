package main

// This is the MountainArray's API interface.
type MountainArray1095 struct {
	arr []int
}

func (this *MountainArray1095) get(index int) int {
	return this.arr[index]
}
func (this *MountainArray1095) length() int {
	return len(this.arr)
}

func findInMountainArray1095(target int, mountainArr *MountainArray1095) int {
	peakIdx := findPeak1095(mountainArr)
	l, r := 0, peakIdx
	for l <= r {
		mid := l + (r-l)/2
		midVal := mountainArr.get(mid)
		if midVal > target {
			r = mid - 1
		} else if midVal < target {
			l = mid + 1
		} else {
			return mid
		}
	}

	l, r = peakIdx, mountainArr.length()-1
	for l <= r {
		mid := l + (r-l)/2
		midVal := mountainArr.get(mid)
		if midVal > target {
			l = mid + 1
		} else if midVal < target {
			r = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func findPeak1095(arr *MountainArray1095) int {
	l, r := 0, arr.length()-1
	for l < r-1 {
		mid := l + (r-l)/2
		midVal := arr.get(mid)
		if arr.get(mid-1) > midVal {
			r = mid
		} else if arr.get(mid+1) > midVal {
			l = mid
		} else {
			return mid
		}
	}
	if arr.get(l) > arr.get(r) {
		return l
	} else {
		return r
	}
}
