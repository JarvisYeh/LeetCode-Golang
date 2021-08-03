package main

import (
	"fmt"
)

func reverseBits190I(num uint32) uint32 {
	i, j := 0, 31
	for i < j {
		one := (num >> i) & 1
		two := (num >> j) & 1
		if one != two {
			num = num ^ (1 << i) ^ (1 << j)
		}
		i++
		j--
	}
	return num
}

func reverseBits190II(num uint32) uint32 {
	res := uint32(0)
	for i := 0; i < 32; i++ {
		res = (res << 1) + ((num >> i) & 1)
	}
	return res
}

func main() {
	i := uint32(17)
	fmt.Printf("%b\n", i)
	i = reverseBits190II(i)
	fmt.Printf("%b\n", i)
}
