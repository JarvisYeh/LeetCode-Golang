package main

import (
	"fmt"
	"strconv"
)

func findLadders126(beginWord string, endWord string, wordList []string) [][]string {
	dict := map[string]bool{}
	for _, w := range wordList {
		dict[w] = true
	}

	// key = s
	// value = list of string
	// list[0] is the minial steps requried from beginWord to s
	// list[1:] is the words that could be reached from beginWords with min steps (e.g. lsit[0])
	prevWordsMap := map[string][]string{
		beginWord: []string{"0"},
	}

	q := []string{beginWord}
	for len(q) != 0 {
		currWord := q[0]
		arr := []byte(currWord)
		currWordsList := prevWordsMap[q[0]]
		currStep := 0
		if s, err := strconv.Atoi(currWordsList[0]); err == nil {
			currStep = s
		}
		q = q[1:]

		// check neighbors
		for i := 0; i < len(arr); i++ {
			tmp := arr[i]
			for c := byte('a'); c <= 'z'; c++ {
				// skip the neighbor if it doesn't chagne any letter
				// or it's not in dict
				if c == tmp {
					continue
				}
				arr[i] = c
				nei := string(arr)
				if nei == "pearl" {
					fmt.Println(nei)
				}
				if _, ok := dict[nei]; !ok {
					continue
				}

				if list, ok := prevWordsMap[nei]; ok {
					if s, err := strconv.Atoi(list[0]); err == nil {
						if currStep+1 == s {
							prevWordsMap[nei] = append(prevWordsMap[nei], currWord)
						}
					}
				} else {
					prevWordsMap[nei] = []string{strconv.Itoa(currStep + 1), currWord}
					q = append(q, nei)
				}

			}
			arr[i] = tmp
		}
	}

	res := [][]string{}
	dfs126(endWord, beginWord, prevWordsMap, []string{}, &res)
	// reverse result list to make it begin -> end
	i, j := 0, len(res)-1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}

func dfs126(currWord, endWord string, stepMap map[string][]string, curr []string, res *[][]string) {
	curr = append(curr, currWord)
	if currWord == endWord {
		tmp := make([]string, len(curr))
		copy(tmp, curr)
		*res = append(*res, tmp)
		return
	}

	if list, ok := stepMap[currWord]; ok {
		for i, word := range list {
			if i == 0 {
				continue
			}
			dfs126(word, endWord, stepMap, curr, res)
		}
	}
	curr = curr[:len(curr)-1]
}

func main() {
	findLadders126("magic", "pearl", []string{"flail", "halon", "lexus", "joint", "pears", "slabs", "lorie", "lapse", "wroth", "yalow", "swear", "cavil", "piety", "yogis", "dhaka", "laxer", "tatum", "provo", "truss", "tends", "deana", "dried", "hutch", "basho", "flyby", "miler", "fries", "floes", "lingo", "wider", "scary", "marks", "perry", "igloo", "melts", "lanny", "satan", "foamy", "perks", "denim", "plugs", "cloak", "cyril", "women", "issue", "rocky", "marry", "trash", "merry", "topic", "hicks", "dicky", "prado", "casio", "lapel", "diane", "serer", "paige", "parry", "elope", "balds", "dated", "copra", "earth", "marty", "slake", "balms", "daryl", "loves", "civet", "sweat", "daley", "touch", "maria", "dacca", "muggy", "chore", "felix", "ogled", "acids", "terse", "cults", "darla", "snubs", "boats", "recta", "cohan", "purse", "joist", "grosz", "sheri", "steam", "manic", "luisa", "gluts", "spits", "boxer", "abner", "cooke", "scowl", "kenya", "hasps", "roger", "edwin", "black", "terns", "folks", "demur", "dingo", "party", "brian", "numbs", "forgo", "gunny", "waled", "bucks", "titan", "ruffs", "pizza", "ravel", "poole", "suits", "stoic", "segre", "white", "lemur", "belts", "scums", "parks", "gusts", "ozark", "umped", "heard", "lorna", "emile", "orbit", "onset", "cruet", "amiss", "fumed", "gelds", "italy", "rakes", "loxed", "kilts", "mania", "tombs", "gaped", "merge", "molar", "smith", "tangs", "misty", "wefts", "yawns", "smile", "scuff", "width", "paris", "coded", "sodom", "shits", "benny", "pudgy", "mayer", "peary", "curve", "tulsa", "ramos", "thick", "dogie", "gourd", "strop", "ahmad", "clove", "tract", "calyx", "maris", "wants", "lipid", "pearl", "maybe", "banjo", "south", "blend", "diana", "lanai", "waged", "shari", "magic", "duchy", "decca", "wried", "maine", "nutty", "turns", "satyr", "holds", "finks", "twits", "peaks", "teems", "peace", "melon", "czars", "robby", "tabby", "shove", "minty", "marta", "dregs", "lacks", "casts", "aruba", "stall", "nurse", "jewry", "knuth"})
}
