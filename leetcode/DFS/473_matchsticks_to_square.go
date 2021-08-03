package main

func makesquare473(matchsticks []int) bool {
	groups := make([]int, 4)
	sum := 0
	for _, val := range matchsticks {
		sum += val
	}
	if sum%4 != 0 {
		return false
	}
	return findMatch473(0, sum/4, matchsticks, groups)
}

func findMatch473(curr int, target int, sticks, groups []int) bool {
	if curr == len(sticks) {
		return groups[0] == groups[1] && groups[1] == groups[2] && groups[2] == groups[3]
	}

	for i := 0; i < len(groups); i++ {
		// prune
		if groups[i]+sticks[curr] > target {
			continue
		}
		// dfs
		groups[i] += sticks[curr]
		if findMatch473(curr+1, target, sticks, groups) {
			return true
		}
		groups[i] -= sticks[curr]
	}
	return false

}
