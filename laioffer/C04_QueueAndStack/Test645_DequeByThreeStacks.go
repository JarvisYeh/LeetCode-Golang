package main

import (
	"fmt"
	"strconv"
)

type Deque645 struct {
	left  []int
	right []int
	buf   []int
}

func Constructor645() Deque645 {
	return Deque645{[]int{}, []int{}, []int{}}
}

func (this *Deque645) OfferFirst(val int) {
	this.left = append(this.left, val)
}

func (this *Deque645) OfferLast(val int) {
	this.right = append(this.right, val)
}

func (this *Deque645) PeekFirst() int {
	if this.Empty() {
		return -1
	}
	if len(this.left) == 0 {
		// shuffle half the elements from right stack to left stack
		this.shuffleDeque(&this.left, &this.right)
	}
	return this.left[len(this.left)-1]
}

func (this *Deque645) PeekLast() int {
	if this.Empty() {
		return -1
	}
	if len(this.right) == 0 {
		// shuffle half the elements from left stack to right stack
		this.shuffleDeque(&this.right, &this.left)
	}
	return this.right[len(this.right)-1]
}

func (this *Deque645) PollFirst() int {
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

func (this *Deque645) PollLast() int {
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

func (this *Deque645) Size() int {
	return len(this.left) + len(this.right)
}

func (this *Deque645) Empty() bool {
	return this.Size() == 0
}

func (this *Deque645) shuffleDeque(dst *[]int, src *[]int) {
	if len(*dst) == 0 {
		// shuffle half from src to buffer
		for i := 0; i < len(*src)/2; i++ {
			num := (*src)[len(*src)-1]
			*src = (*src)[:len(*src)-1]
			this.buf = append(this.buf, num)
		}

		// shuffle the other half from src to dst stack
		for len(*src) > 0 {
			num := (*src)[len(*src)-1]
			*src = (*src)[:len(*src)-1]
			*dst = append(*dst, num)
		}

		// shuffle the buf back to src stack
		for len(this.buf) > 0 {
			num := this.buf[len(this.buf)-1]
			this.buf = this.buf[:len(this.buf)-1]
			*src = append(*src, num)
		}
	}
}

func (this *Deque645) printDeque() {
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
	deque := Constructor645()
	for _, i := range []int{1, 2, 3, 4, 5, 6, 7, 8} {
		deque.OfferFirst(i)
		deque.OfferLast(2 * i)
	}
	deque.printDeque()

	for i := 0; i < 11; i++ {
		deque.PollFirst()
		deque.printDeque()
	}
	for i := 0; i < 4; i++ {
		deque.PollLast()
		deque.printDeque()
	}
}
