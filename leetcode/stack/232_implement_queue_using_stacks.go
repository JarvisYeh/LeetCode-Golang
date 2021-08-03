package main

type MyQueue232 struct {
	in  *[]int
	out *[]int
}

/** Initialize your data structure here. */
func Constructor232() MyQueue232 {
	return MyQueue232{&[]int{}, &[]int{}}
}

/** Push element x to the back of queue. */
func (this *MyQueue232) Push(x int) {
	*this.in = append(*this.in, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue232) Pop() int {
	if len(*this.out) == 0 {
		this.shuffle()
	}
	if len(*this.out) > 0 {
		poped := (*this.out)[len(*this.out)-1]
		*this.out = (*this.out)[:len(*this.out)-1]
		return poped
	} else {
		return -1
	}
}

/** Get the front element. */
func (this *MyQueue232) Peek() int {
	if len(*this.out) == 0 {
		this.shuffle()
	}
	if len(*this.out) > 0 {
		return (*this.out)[len(*this.out)-1]
	} else {
		return -1
	}
}

/** Returns whether the queue is empty. */
func (this *MyQueue232) Empty() bool {
	return (len(*this.in) + len(*this.out)) == 0
}

func (this *MyQueue232) shuffle() {
	for len(*this.in) > 0 {
		poped := (*this.in)[len(*this.in)-1]
		*this.in = (*this.in)[:len(*this.in)-1]
		*this.out = append(*this.out, poped)
	}
}
