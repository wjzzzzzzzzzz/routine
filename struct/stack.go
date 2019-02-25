package Struct

import (
	"sync"
)

type Stack struct {
	size  int
	top   int
	value []int
	sync.RWMutex
}

func (s *Stack) Pop() (rtn int, success bool) {
	if !s.IsNull() {
		s.top--
		rtn = s.value[s.top]

		return rtn, true
	}
	return rtn, false
}

func (s *Stack) Push(a int) bool {
	if !s.IsFull() {
		s.value[s.top] = a
		s.top++
		return true
	}
	return false
}

func (s *Stack) Pushs(a ...int) bool {
	if !s.IsFull() && s.size-s.top >= len(a) {
		for _, v := range a {
			s.Push(v)
		}
	}
	return false

}

func (s *Stack) IsNull() bool {

	if s.top == -1 {
		return true
	} else {
		return false
	}

}
func (s *Stack) IsFull() bool {

	if s.size <= s.top {
		return true
	} else {
		return false
	}

}
func (s *Stack) Show() []int {

	return s.value

}
func NewStack(size int) *Stack {

	return &Stack{
		size:  size,
		top:   0,
		value: make([]int, size),
	}

}
