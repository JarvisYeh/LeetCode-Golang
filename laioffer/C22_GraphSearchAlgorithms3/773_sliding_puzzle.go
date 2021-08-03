package main

// TC: O(n(V + E))
// SC: O(n*V)
func slidingPuzzle773(board [][]int) int {
	m, n := len(board), len(board[0])
	start, end := make([]byte, m*n), "123450"
	for i, _ := range start {
		start[i] = byte(board[i/n][i%n] + '0')
	}

	q := []string{string(start)}
	generated := map[string]bool{
		string(start): true,
	}

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	step := 0
	for len(q) != 0 {
		// layer by layer, each layer has size = len(q)
		size := len(q)
		for t := 0; t < size; t++ {
			curr := q[0]
			q = q[1:]
			if curr == end {
				return step
			}
			zeroIndex := findZero(curr)
			for _, dir := range dirs {
				i, j := zeroIndex/n+dir[0], zeroIndex%n+dir[1]
				if i < 0 || i >= m || j < 0 || j >= n {
					continue
				}
				newBoard := []byte(curr)
				// swap zero with neighbor positions
				newBoard[zeroIndex], newBoard[i*n+j] = newBoard[i*n+j], newBoard[zeroIndex]
				if _, ok := generated[string(newBoard)]; !ok {
					q = append(q, string(newBoard))
					generated[string(newBoard)] = true
				}
			}

		}
		step++
	}
	return -1
}

func findZero(board string) int {
	for i := 0; i < len(board); i++ {
		if board[i] == '0' {
			return i
		}
	}
	return -1
}
