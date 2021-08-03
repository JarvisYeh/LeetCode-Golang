package main

import "fmt"

type TrieNode212 struct {
	isWord   bool
	children []*TrieNode212
}

type Trie212 struct {
	root *TrieNode212
}

func (t *Trie212) insert(word string) bool {
	curr := t.root
	for i := 0; i < len(word); i++ {
		if curr.children[word[i]-'a'] == nil {
			curr.children[word[i]-'a'] = &TrieNode212{false, make([]*TrieNode212, 26)}
		}
		curr = curr.children[word[i]-'a']
	}

	// return if insert successfully
	if curr.isWord {
		return false
	} else {
		curr.isWord = true
		return true
	}
}

func buildTrie(words []string) *TrieNode212 {
	trie := Trie212{nil}
	trie.root = &TrieNode212{false, make([]*TrieNode212, 26)}
	for _, word := range words {
		trie.insert(word)
	}
	return trie.root
}

// K words in dict, each with length L
// TC: O(K*L + m*n*4^L)
// SC: O(m*n + L + Trie(26^L))
func findWords212(board [][]byte, words []string) []string {
	m, n := len(board), len(board[0])
	set := map[string]bool{}

	// O(K*L) to insert K words with length L into Trie
	root := buildTrie(words)

	// for each start position
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if root.children[board[i][j]-'a'] != nil {
				visited := make([][]bool, m)
				for i, _ := range visited {
					visited[i] = make([]bool, n)
				}
				visited[i][j] = true
				helper212(board, i, j, root.children[board[i][j]-'a'], visited,
					[]byte{board[i][j]}, set)
			}
		}
	}

	res := []string{}
	for key, _ := range set {
		res = append(res, key)
	}
	return res
}

// visited are the positions from start position to (i, j), include (i, j)
// curr is node with prefix letters form start position to (i, j), include (i, j)
// currPath are all letters corresponds to curr TrieNode
func helper212(board [][]byte, i, j int, curr *TrieNode212, visited [][]bool, currPath []byte, set map[string]bool) {
	m, n := len(board), len(board[0])

	// if all letters from start position to (i, j) form a word
	// append to result
	if curr.isWord {
		set[string(currPath)] = true
	}

	dirs := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	for _, dir := range dirs {
		newI := i + dir[0]
		newJ := j + dir[1]
		if newI < 0 || newI >= m || newJ < 0 || newJ >= n || visited[newI][newJ] {
			continue
		}
		child := curr.children[board[newI][newJ]-'a']
		if child != nil {
			visited[newI][newJ] = true
			currPath = append(currPath, board[newI][newJ])
			helper212(board, newI, newJ, child, visited, currPath, set)
			currPath = currPath[:len(currPath)-1]
			visited[newI][newJ] = false
		}
	}
	// there is no explicit base case here
	// since there is condition if child != nil, then dfs(), so it won't reach the leave node
}

func main() {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}

	dict := []string{"oath", "eat"}

	fmt.Println(findWords212(board, dict))
}
