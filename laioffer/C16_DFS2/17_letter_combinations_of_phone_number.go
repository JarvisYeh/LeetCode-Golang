package main

//DFS
func letterCombinations17I(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	keyMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	res := &[]string{}
	curr := []byte{}
	helper17(0, digits, keyMap, curr, res)
	return *res
}

func helper17(idx int, digits string, keyMap map[byte]string, curr []byte, res *[]string) {
	if idx == len(digits) {
		*res = append(*res, string(curr))
		return
	}

	chars := keyMap[digits[idx]]
	for i := 0; i < len(chars); i++ {
		helper17(idx+1, digits, keyMap, append(curr, chars[i]), res)
	}
}

// BFS
func letterCombinations17II(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	keyMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	queue := [][]byte{}
	queue = append(queue, []byte{})
	for i := 0; i < len(digits); i++ {
		chars := keyMap[digits[i]]
		size := len(queue)
		for j := 0; j < size; j++ {
			poped := queue[0]
			tmp := make([]byte, len(poped))
			copy(tmp, poped)
			queue = queue[1:]
			for k := 0; k < len(chars); k++ {
				queue = append(queue, append(tmp, chars[k]))
			}
		}
	}
	res := []string{}
	for _, val := range queue {
		res = append(res, string(val))
	}
	return res
}
