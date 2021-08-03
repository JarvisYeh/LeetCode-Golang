package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1.0 / powPositive(x, -n)
	} else if n > 0 {
		return powPositive(x, n)
	} else {
		return 1.0
	}
}

func powPositive(x float64, n int) float64 {
	pow := 1.0
	remain := n
	for remain > 0 {
		currentPow := x
		count := 1
		for (count << 1) <= remain {
			count <<= 1
			currentPow *= currentPow
		}
		pow *= currentPow
		remain -= count
	}
	return pow
}

func main() {
	fmt.Println(myPow(2.0, 5))
}
