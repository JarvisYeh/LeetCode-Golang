package main

import (
	"fmt"
	"sort"
)

// unsorted array with duplication

// method 1
// use two hashmap to store frequency of two arrays
// TC: O(m + n)
// SC: O(m + n)
func common651I(a, b []int) []int {
	mapA, mapB := make(map[int]int), make(map[int]int)

	for _, num := range a {
		if count, ok := mapA[num]; ok {
			mapA[num] = count + 1
		} else {
			mapA[num] = 1
		}
	}

	for _, num := range b {
		if count, ok := mapB[num]; ok {
			mapB[num] = count + 1
		} else {
			mapB[num] = 1
		}
	}
	res := []int{}
	// make mapA the smaller one, iterate the smaller one
	if len(mapA) > len(mapB) {
		mapA, mapB = mapB, mapA
	}
	for num, countA := range mapA {
		if countB, ok := mapB[num]; ok {
			if countB < countA {
				countA = countB
			}
			for i := 0; i < countA; i++ {
				res = append(res, num)
			}
		}
	}
	return res
}

// method 2
// sort then use two pointers
// TC: O(nlogn + mlogm)
// SC: O(1) if use heap sort
func common651II(a, b []int) []int {
	sort.Ints(a)
	sort.Ints(b)
	i, j := 0, 0
	res := []int{}
	for i < len(a) && j < len(b) {
		if a[i] > b[j] {
			j++
		} else if a[i] < b[j] {
			i++
		} else {
			res = append(res, a[i])
			j++
			i++
		}
	}
	return res
}

func main() {
	fmt.Println(common651II(
		[]int{1, 9, -5, 6, 7, 2, 2, 100, 100},
		[]int{9, 4, 100, 2, 9, 9, 100}))
}
