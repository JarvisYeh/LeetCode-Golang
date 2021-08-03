package main

func restoreIpAddresses93(s string) []string {
	res := &[]string{}
	dfs93(s, 0, 0, []byte{}, res)
	return *res
}

func dfs93(s string, idx int, block int, curr []byte, res *[]string) {
	// base case
	if idx == len(s) && block == 4 {
		*res = append(*res, string(curr[:len(curr)-1]))
		return
	}

	if idx == len(s) || block == 4 {
		return
	}

	// one block can have at most 3 digits
	count := 0
	for i := idx; i < idx+3; i++ {
		if i >= len(s) {
			return
		}
		// if current first digits, which is s[idx] is '0', only one possible next level dfs
		if s[idx] == '0' {
			dfs93(s, idx+1, block+1, append(curr, '0', '.'), res)
			break
		}
		count = count*10 + int(s[i]-'0')
		if count <= 255 {
			curr = append(curr, s[i], '.')
			dfs93(s, i+1, block+1, curr, res)
			curr = curr[:len(curr)-1]
		}
	}

}

func main() {
	restoreIpAddresses93("010010")
}
