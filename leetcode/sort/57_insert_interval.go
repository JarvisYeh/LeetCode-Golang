package main

func insert57(intervals [][]int, newInterval []int) [][]int {
	if intervals == nil || len(intervals) == 0 {
		return [][]int{newInterval}
	}
	res := make([][]int, 0)
	i := 0
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}

	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		if newInterval[0] > intervals[i][0] {
			newInterval[0] = intervals[i][0]
		}
		if newInterval[1] < intervals[i][1] {
			newInterval[1] = intervals[i][1]
		}
		i++
	}
	res = append(res, newInterval)

	for i < len(intervals) {
		res = append(res, intervals[i])
		i++
	}
	return res
}
