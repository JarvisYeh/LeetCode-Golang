package main

type Node133 struct {
	Val       int
	Neighbors []*Node133
}

// BFS
//
func cloneGraph133BFS(node *Node133) *Node133 {
	if node == nil {
		return node
	}

	q := []*Node133{}
	generated := make(map[*Node133]*Node133)

	// init
	q = append(q, node)
	generated[node] = &Node133{node.Val, []*Node133{}}

	// loop
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		mapNode := generated[curr]
		// generate each neighbors who are not generated before
		for _, nei := range curr.Neighbors {
			if _, ok := generated[nei]; !ok {
				tmp := &Node133{nei.Val, []*Node133{}}
				// generate the neighbors nodes that are not generate
				q = append(q, nei)
				generated[nei] = tmp
			}
			mapNode.Neighbors = append(mapNode.Neighbors, generated[nei])
		}

	}
	return generated[node]
}

// DFS
func cloneGraph133DFS(node *Node133) *Node133 {
	if node == nil {
		return nil
	}
	clonedNodes := make(map[*Node133]*Node133)
	return dfs133(node, clonedNodes)
}

func dfs133(node *Node133, clonedNodes map[*Node133]*Node133) *Node133 {
	// base case, node is already cloned (with all it neighbors info)
	if n, ok := clonedNodes[node]; ok {
		return n
	}

	curr := &Node133{node.Val, []*Node133{}}
	clonedNodes[node] = curr
	for _, nei := range node.Neighbors {
		curr.Neighbors = append(curr.Neighbors, dfs133(nei, clonedNodes))
	}
	return curr
}
