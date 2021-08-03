package main

import (
	"fmt"
)

type ArrayReader702 struct {
	arr []int
}

func (this *ArrayReader702) get(index int) int {
	if index >= len(this.arr) {
		return 1<<31 - 1
	} else {
		return this.arr[index]
	}
}

func constructor702(arr []int) ArrayReader702 {
	reader := ArrayReader702{}
	reader.arr = arr
	return reader
}

// search in an unknown size array
// reader.get(idx)
// return 2^32 - 1 if out of boundary
// return array[idx] if in boundary
func search(reader ArrayReader702, target int) int {
	r := 1
	for reader.get(r) >= target {
		r *= 2
	}

	l := r / 2
	for l <= r {
		mid := l + (r-l)/2
		if reader.get(mid) == target {
			return mid
		} else if reader.get(mid) > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func main() {
	reader := constructor702([]int{-1, 0, 3, 5, 9, 12})
	fmt.Println(search(reader, 5))
}
