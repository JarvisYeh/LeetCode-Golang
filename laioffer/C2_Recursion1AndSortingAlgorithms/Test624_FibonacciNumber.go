package main

// recursion
func fibo624I(K int) int {
	if K == 0 {
		return 0
	}
	if K == 1 {
		return 1
	}

	return fibo624I(K-1) + fibo624I(K-2)
}

// iteration
func fibo624II(K int) int {
	if K == 0 {
		return 0
	}
	if K == 1 {
		return 1
	}

	a, b := 0, 1
	for i := 2; i <= K; i++ {
		tmp := b
		b = a + b
		a = tmp
	}
	return b
}
