package main

import (
	"fmt"
	"strconv"
)

type DeuqeByTwoStacks struct {
	left  []int
	right []int
}

func ConstructorDequeByTwoStacks() DeuqeByTwoStacks {
	return DeuqeByTwoStacks{[]int{}, []int{}}
}

func (this *DeuqeByTwoStacks) OfferFirst(val int) {
	this.left = append(this.left, val)
}

func (this *DeuqeByTwoStacks) OfferLast(val int) {
	this.right = append(this.right, val)
}

func (this *DeuqeByTwoStacks) PeekFirst() int {
	if this.Empty() {
		return -1
	}
	if len(this.left) == 0 {
		// shuffle half the elements from right stack to left stack
		this.shuffleDeque(&this.left, &this.right)
	}
	return this.left[len(this.left)-1]
}

func (this *DeuqeByTwoStacks) PeekLast() int {
	if this.Empty() {
		return -1
	}
	if len(this.right) == 0 {
		// shuffle half the elements from left stack to right stack
		this.shuffleDeque(&this.right, &this.left)
	}
	return this.right[len(this.right)-1]
}

func (this *DeuqeByTwoStacks) PollFirst() int {
	if this.Empty() {
		return -1
	}
	if len(this.left) == 0 {
		// shuffle half the elements from right stack to left stack
		this.shuffleDeque(&this.left, &this.right)
	}
	popped := this.left[len(this.left)-1]
	this.left = this.left[:len(this.left)-1]
	return popped
}

func (this *DeuqeByTwoStacks) PollLast() int {
	if this.Empty() {
		return -1
	}
	if len(this.right) == 0 {
		// shuffle half the elements from left stack to right stack
		this.shuffleDeque(&this.right, &this.left)
	}
	popped := this.right[len(this.right)-1]
	this.right = this.right[:len(this.right)-1]
	return popped
}

func (this *DeuqeByTwoStacks) Size() int {
	return len(this.left) + len(this.right)
}

func (this *DeuqeByTwoStacks) Empty() bool {
	return this.Size() == 0
}

func (this *DeuqeByTwoStacks) shuffleDeque(dst *[]int, src *[]int) {
	if len(*dst) == 0 {
		// shuffle all elements in src to dst
		for len(*src) > 0 {
			num := (*src)[len(*src)-1]
			*src = (*src)[:len(*src)-1]
			*dst = append(*dst, num)
		}
	}
}

func (this *DeuqeByTwoStacks) printDeque() {
	for i := len(this.left) - 1; i >= 0; i-- {
		fmt.Print(strconv.Itoa(this.left[i]) + "->")
	}
	fmt.Print("|")
	for _, val := range this.right {
		fmt.Print(strconv.Itoa(val) + "->")
	}
	fmt.Println()
}

func main() {
	deque := ConstructorDequeByTwoStacks()
	for _, i := range []int{1, 2, 3, 4, 5, 6, 7, 8} {
		deque.OfferFirst(i)
		deque.OfferLast(2 * i)
	}
	deque.printDeque()

	for i := 0; i < 9; i++ {
		deque.PollFirst()
		deque.printDeque()
	}

	for i := 0; i < 3; i++ {
		deque.PeekLast()
		deque.printDeque()
		deque.PeekFirst()
		deque.printDeque()
	}
}
