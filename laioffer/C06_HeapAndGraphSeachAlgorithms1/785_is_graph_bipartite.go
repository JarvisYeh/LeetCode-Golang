package main

// BFS
func isBipartite785BFS(graph [][]int) bool {
	groups := make([]int, len(graph))
	for i := 0; i < len(graph); i++ {
		if !checkValid785BFS(i, graph, groups) {
			return false
		}
	}
	return true
}

func checkValid785BFS(graph int, neighbors [][]int, groups []int) bool {
	// this individual graph already checked before, return true
	if groups[graph] != 0 {
		return true
	}

	// initialization
	groups[graph] = 1
	queue := []int{}
	queue = append(queue, graph)

	for len(queue) > 0 {
		// expand one node from queue
		curr := queue[0]
		queue = queue[1:]
		currGroup := groups[curr]

		// generate curr's neighbors
		for _, nei := range neighbors[curr] {
			if groups[nei] == 0 {
				groups[nei] = -currGroup
				queue = append(queue, nei)
			} else if groups[nei] == currGroup {
				return false
			}
		}
	}
	return true
}

// DFS
func isBipartite785DFS(graph [][]int) bool {
	groups := make([]int, len(graph))
	for i := 0; i < len(graph); i++ {
		if groups[i] == 0 && !validCheck785DFS(i, 1, groups, graph) {
			return false
		}
	}
	return true
}

func validCheck785DFS(curr, supposedGroup int, groups []int, graph [][]int) bool {
	// base case 1
	if groups[curr] == -supposedGroup {
		return false
	}
	// base case 2
	if groups[curr] == supposedGroup {
		return true
	}

	// traverse that node
	if groups[curr] == 0 {
		groups[curr] = supposedGroup
	}

	// check all its neighbors
	for _, nei := range graph[curr] {
		if !validCheck785DFS(nei, -supposedGroup, groups, graph) {
			return false
		}
	}
	return true
}

func main() {
	isBipartite785BFS([][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}})
}
