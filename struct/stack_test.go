package Struct

import (
	"testing"
	"log"
)

func Test_stack(t *testing.T) {
	s := NewStack(5)
	s.Pushs(1,2,3,4,5)
	log.Println(s.Show())
	for i:=0;i<5 ;i++  {
		if a,ok:=s.Pop();ok{
			log.Println(a)
		}

	}

}
func Test_addTwoNumbers1(t *testing.T) {
	ints := make([]int, 5)

	log.Println(ints)
}