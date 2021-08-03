package main

type MinStack155 struct {
	stack *[]int
	min   *[]int
}

/** initialize your data structure here. */
func Constructor155() MinStack155 {
	return MinStack155{&[]int{}, &[]int{}}
}

func (this *MinStack155) Push(val int) {
	if len(*this.min) == 0 || val <= (*this.min)[len(*this.min)-1] {
		*this.min = append(*this.min, val)
	}
	*this.stack = append(*this.stack, val)
}

func (this *MinStack155) Pop() {
	if len(*this.min) == 0 {
		return
	}
	if (*this.min)[len(*this.min)-1] == (*this.stack)[len(*this.stack)-1] {
		*this.min = (*this.min)[:len(*this.min)-1]
	}
	*this.stack = (*this.stack)[:len(*this.stack)-1]
}

func (this *MinStack155) Top() int {
	if len(*this.stack) == 0 {
		return -1
	}
	return (*this.stack)[len(*this.stack)-1]
}

func (this *MinStack155) GetMin() int {
	if len(*this.min) == 0 {
		return -1
	}
	return (*this.min)[len(*this.min)-1]
}
