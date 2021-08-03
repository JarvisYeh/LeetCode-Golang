package main

import "fmt"

type MyStack225 struct {
	in  []int
	buf []int
}

/** Initialize your data structure here. */
func Constructor225() MyStack225 {
	return MyStack225{[]int{}, []int{}}
}

/** Push element x onto stack. */
func (this *MyStack225) Push(x int) {
	this.in = append(this.in, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack225) Pop() int {
	if len(this.in) == 0 {
		return -1
	}

	for i := len(this.in); i > 1; i-- {
		this.buf = append(this.buf, this.in[0])
		this.in = this.in[1:]
	}
	poped := this.in[0]
	this.in = this.in[1:]
	this.in, this.buf = this.buf, this.in
	return poped
}

/** Get the top element. */
func (this *MyStack225) Top() int {
	if len(this.in) == 0 {
		return -1
	}

	num := -1
	for i := len(this.in); i > 0; i-- {
		num = this.in[0]
		this.buf = append(this.buf, num)
		this.in = this.in[1:]
	}
	this.in, this.buf = this.buf, this.in
	return num
}

/** Returns whether the stack is empty. */
func (this *MyStack225) Empty() bool {
	return len(this.in) == 0
}

func main() {
	stack := Constructor225()
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Top())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Empty())
}
