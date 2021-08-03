package main

func hammingDistance461(x int, y int) int {
	xor := x ^ y
	// use recursion to count bits
	return countBits461(xor)
}

func countBits461(num int) int {
	if num == 0 {
		return 0
	}
	return countBits461(num>>1) + (num & 1)
}

func main() {
	hammingDistance461(1, 4)
}
