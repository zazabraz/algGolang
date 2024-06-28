package main

import "fmt"

type MinStack struct {
	st  []int
	min []int
}

func (this *MinStack) Info() {
	fmt.Printf("st:%v \n min: %v \n", this.st, this.min)
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if len(this.st) == 0 && len(this.min) == 0 {
		this.min = append(this.min, val)
	} else if val <= this.min[len(this.min)-1] {
		this.min = append(this.min, val)
	}

	this.st = append(this.st, val)

}

func (this *MinStack) Pop() {
	if len(this.st) == 0 {
		return
	}

	if this.st[len(this.st)-1] == this.min[len(this.min)-1] {
		this.min = this.min[:len(this.min)-1]
	}

	this.st = this.st[:len(this.st)-1]
}

func (this *MinStack) Top() int {
	if len(this.st) == 0 {
		return 0
	}
	return this.st[len(this.st)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.min) == 0 {
		return 0
	}
	return this.min[len(this.min)-1]
}

func main() {
	st := Constructor()
	st.Push(2)
	st.Push(0)
	st.Push(3)
	st.Push(0)

	st.Info()

	fmt.Println(st.GetMin())
	st.Pop()
	st.Info()

	fmt.Println(st.GetMin())
	st.Pop()
	st.Info()

	fmt.Println(st.GetMin())
	st.Pop()
	st.Info()

	fmt.Println(st.GetMin())
}
