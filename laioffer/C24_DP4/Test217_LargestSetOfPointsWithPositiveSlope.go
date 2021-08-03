package main

import (
	"fmt"
	"sort"
)

type Point217 struct {
	X int
	Y int
}

func largest217(points []Point217) int {
	// sort input w.r.t. X
	sort.Slice(points, func(i, j int) bool {
		return points[i].X < points[j].X
	})
	DP := make([]int, len(points))
	max := 1
	for i := 0; i < len(points); i++ {
		DP[i] = 1
		for j := i - 1; j >= 0; j-- {
			// i.X != j.X otherwise it's slope = infinity
			if points[j].Y < points[i].Y && points[j].X != points[i].X && DP[j]+1 > DP[i] {
				DP[i] = DP[j] + 1
			}
		}
		if max < DP[i] {
			max = DP[i]
		}
	}

	// if only 1 point in result set, return 0
	if max == 1 {
		return 0
	}
	return max
}

func main() {
	fmt.Println(largest217([]Point217{
		Point217{1, 5},
		Point217{1, 3},
		Point217{0, 10},
	}))
}
