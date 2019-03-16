package Struct

import (
	"testing"
	"log"
	"time"
)

func Test_stack(t *testing.T) {
	stack := NewStack(10)

	for i := 0; i < 10; i++ {

		go stack.Push(i)

	}
	log.Println(stack.Show())

	for {
		time.Sleep(time.Second)
	}

}
func Test_addTwoNumbers1(t *testing.T) {
	ints := make([]int, 5)

	log.Println(ints)
}
func Test_slice(t *testing.T) {
	ints := make([]int, 6, 8)
	ints[1] = 8
	log.Println(ints)
	fn := func(a []int, b int) {
		a[1] = b
	}
	fn(ints, 5)
	log.Println(ints)
	fn(ints, 6)
	log.Println(ints)

}
func Test_MinStack(t *testing.T) {
	minStack := Constructor()
	minStack.Push(5)
	minStack.Push(3)
	minStack.Push(1)
	log.Printf("%+v", minStack)
	minStack.Pop()
	log.Println(minStack.GetMin())

	log.Printf("%+v", minStack)
	//
	//log.Println(minStack.GetMin())

}
func Test_NewList(t *testing.T) {
	lists := NewList(1)
	log.Println(lists)
	for {
		log.Println(lists.Val)
		if lists.Next != nil {
			lists = lists.Next
		} else {
			break
		}

	}

}
