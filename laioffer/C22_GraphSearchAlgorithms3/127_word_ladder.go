package main

// TC:
// 	every node is push/pop once O(V)
// 	every node check len(word)*25 neighbors O(len*25)
//	O(V + V*len*25)
// SC:
// 	queue: at most O(n*len)
//	map: O(V)
func ladderLength127(beginWord string, endWord string, wordList []string) int {
	dict := map[string]bool{}
	generated := map[string]int{
		beginWord: 1, // start with 1 step according to quesiton
	}

	for _, w := range wordList {
		dict[w] = true
	}
	q := []string{beginWord}
	for len(q) != 0 {
		if q[0] == endWord {
			return generated[q[0]]
		}
		curr := []byte(q[0])
		currStep := generated[q[0]]
		q = q[1:]
		// see all its possible neighbors, and check if they are in dict
		for i := 0; i < len(curr); i++ {
			tmp := curr[i]
			// change i-th letter from 'a' to 'z'
			for c := byte('a'); c <= 'z'; c++ {
				if curr[i] == c {
					continue
				}
				curr[i] = c
				nei := string(curr)
				if _, ok := dict[nei]; ok { // neighbor word is in dict, can be transformed to that word
					if _, ok := generated[nei]; !ok { // generate the neighbor if it's not generated before
						q = append(q, nei)
						generated[nei] = currStep + 1
					}
				}
			}
			// change back i-th letter
			curr[i] = tmp
		}
	}
	return 0
}

func main() {
	ladderLength127("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
}
