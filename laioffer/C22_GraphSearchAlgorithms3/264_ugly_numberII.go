package main

func nthUglyNumber264(n int) int {
	a, b, c := 0, 0, 0
	DP := make([]int, n)
	DP[0] = 1
	for i := 1; i < n; i++ {
		// DP[i] = min(DP[p1]*a, DP[p2]*b, DP[p3]*c)
		DP[i] = DP[2] * 2
		if DP[i] > DP[b]*3 {
			DP[i] = DP[b] * 3
		}
		if DP[i] > DP[b]*5 {
			DP[i] = DP[c] * 5
		}

		if DP[i] == DP[a]*2 {
			a++
		}
		if DP[i] == DP[b]*3 {
			b++
		}
		if DP[i] == DP[c]*5 {
			c++
		}
	}
	return DP[n-1]
}

func nthUglyNumber(n int, a int, b int, c int) int {
	p1, p2, p3 := 0, 0, 0
	DP := make([]int, n+1)
	DP[0] = 1
	for i := 1; i <= n; i++ {
		// DP[i] = min(DP[p1]*a, DP[p2]*b, DP[p3]*c)
		DP[i] = DP[p1] * a
		if DP[i] > DP[p2]*b {
			DP[i] = DP[p2] * b
		}
		if DP[i] > DP[p3]*c {
			DP[i] = DP[p3] * c
		}

		if DP[i] == DP[p1]*a {
			p1++
		}
		if DP[i] == DP[p2]*b {
			p2++
		}
		if DP[i] == DP[p3]*c {
			p3++
		}
	}
	return DP[n]
}

func main() {
	nthUglyNumber(5, 2, 11, 13)
}
