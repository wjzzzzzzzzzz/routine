package Struct


//维护最小值的栈
type MinStack struct {
	min   int
	value []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		min:   0,
		value: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.value = append(this.value, x)
	if len(this.value)==1{
		this.min=x
		return
	}
	if this.min > x {
		this.min = x
	}
}

func (this *MinStack) Pop() {
	if len(this.value)==1{
		this.min=0
		this.value=this.value[:0]
		return
	}
	if this.value[len(this.value)-1] == this.min {
		min := this.value[0]
		for i := 0; i < len(this.value)-1; i++ {
			if min > this.value[i] {
				min = this.value[i]
			}

		}
		this.min = min
	}
	this.value=this.value[:len(this.value)-1]
}

func (this *MinStack) Top() int {
	return this.value[len(this.value)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}
