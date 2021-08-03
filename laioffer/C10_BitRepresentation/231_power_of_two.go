package main

import "fmt"

func isPowerOfTwo231I(n int) bool {
	if n < 0 {
		return false
	}

	count := 0
	for i := 0; i < 31; i++ {
		count += n >> i & 1
	}
	return count == 1
}

func isPowerOfTwo231II(n int) bool {
	return n > 0 && n&(n-1) == 0
}

func main() {
	fmt.Println(isPowerOfTwo231II(256))
}
