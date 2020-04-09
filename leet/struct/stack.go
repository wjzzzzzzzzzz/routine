package Struct

import (
	"routine/leet"
	"sync"
)

type MyStack struct {
	size  int
	top   int
	value []int
	sync.RWMutex
}

func (s *MyStack) Pop() (rtn int, success bool) {
	if !s.IsNull() {
		s.top--
		rtn = s.value[s.top]

		return rtn, true
	}
	return rtn, false
}

func (s *MyStack) Push(a int) bool {
	if !s.IsFull() {
		s.value[s.top] = a
		s.top++
		return true
	}
	return false
}

func (s *MyStack) Pushs(a ...int) bool {
	if !s.IsFull() && s.size-s.top >= len(a) {
		for _, v := range a {
			s.Push(v)
		}
	}
	return false

}

func (s *MyStack) IsNull() bool {

	if s.top == -1 {
		return true
	} else {
		return false
	}

}
func (s *MyStack) IsFull() bool {

	if s.size <= s.top {
		return true
	} else {
		return false
	}

}
func (s *MyStack) Show() []int {

	return s.value

}
func NewStack(size int) main.Stack {
	return &MyStack{
		size:  size,
		top:   0,
		value: make([]int, size),
	}

}
