package main

import "container/heap"

type Cell407 struct {
	i      int
	j      int
	height int
	index  int
}

type MinHeap407 []*Cell407

func (pq MinHeap407) Len() int {
	return len(pq)
}

func (pq MinHeap407) Less(i, j int) bool {
	return pq[i].height < pq[j].height
}

func (pq MinHeap407) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *MinHeap407) Push(x interface{}) {
	ele := x.(*Cell407)
	ele.index = len(*pq)
	*pq = append(*pq, ele)
}

func (pq *MinHeap407) Pop() interface{} {
	old := *pq
	n := len(*pq)
	popped := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]
	return popped
}

func trapRainWater(heightMap [][]int) int {
	m, n := len(heightMap), len(heightMap[0])
	// initialize minHeap and generated matrix
	minHeap := MinHeap407{}
	generated := make([][]bool, m)
	for i, _ := range generated {
		generated[i] = make([]bool, n)
	}

	// left and right sides
	for i := 0; i < m; i++ {
		minHeap = append(minHeap,
			&Cell407{i, 0, heightMap[i][0], len(minHeap)},
			&Cell407{i, n - 1, heightMap[i][n-1], len(minHeap) + 1})
		generated[i][0] = true
		generated[i][n-1] = true
	}

	// top and bottom sides
	for j := 0; j < n; j++ {
		minHeap = append(minHeap,
			&Cell407{0, j, heightMap[0][j], len(minHeap)},
			&Cell407{m - 1, j, heightMap[m-1][j], len(minHeap) + 1})
		generated[0][j] = true
		generated[m-1][j] = true
	}

	vol := 0
	heap.Init(&minHeap)
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(minHeap) != 0 {
		curr := heap.Pop(&minHeap).(*Cell407)
		// check left, right, up, down four neighbours of current cell
		for _, dir := range dirs {
			i := curr.i + dir[0]
			j := curr.j + dir[1]
			if i < 0 || i >= m || j < 0 || j >= n || generated[i][j] {
				continue
			}
			// if neighbor's height is less than popped smallest height
			// which means some water can be trapped
			height := heightMap[i][j]
			if curr.height > height {
				vol += curr.height - height
			}

			// generate that neighbor with height = max(self.height, curr.height
			if curr.height > height {
				height = curr.height
			}
			heap.Push(&minHeap, &Cell407{i, j, height, -1})
			generated[i][j] = true
		}
	}
	return vol
}
