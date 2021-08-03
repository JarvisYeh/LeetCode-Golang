package main

import "fmt"

type Tuple315 struct {
	Val   int
	Count int
}

func countSmaller315(nums []int) []int {
	// cellMap stores <index in nums, Tuple>
	cellMap := make(map[int]*Tuple315)

	// arr is the Tuple array
	arr := make([]*Tuple315, len(nums))

	for i, val := range nums {
		arr[i] = &Tuple315{val, 0}
		cellMap[i] = arr[i]
	}

	mergeSort315(arr, 0, len(arr)-1)

	// after mergeSort, tuple array is sorted w.r.t. its Val
	// Count stores how many elements are moved to left of it
	res := make([]int, len(nums))
	for i := 0; i < len(res); i++ {
		res[i] = cellMap[i].Count
	}
	return res
}

func mergeSort315(arr []*Tuple315, l, r int) []*Tuple315 {
	if l >= r {
		return []*Tuple315{arr[l]}
	}
	mid := l + ((r - l) >> 1)
	left := mergeSort315(arr, l, mid)
	right := mergeSort315(arr, mid+1, r)
	return merge315(left, right)
}

func merge315(lArr, rArr []*Tuple315) []*Tuple315 {
	l, r, res := 0, 0, 0
	resArr := make([]*Tuple315, len(lArr)+len(rArr))
	for l < len(lArr) && r < len(rArr) {
		if lArr[l].Val <= rArr[r].Val {
			resArr[res] = lArr[l]
			resArr[res].Count += r
			res++
			l++
		} else {
			resArr[res] = rArr[r]
			res++
			r++
		}
	}
	for l < len(lArr) {
		resArr[res] = lArr[l]
		resArr[res].Count += r
		l++
		res++
	}
	for r < len(rArr) {
		resArr[res] = rArr[r]
		r++
		res++
	}

	return resArr
}

func main() {
	fmt.Println(countSmaller315([]int{5, 2, 6, 1}))
}
