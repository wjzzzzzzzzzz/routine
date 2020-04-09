package Struct

import "routine/leet"

//

func NewList(val ...int) *main.ListNode {
	var temp *main.ListNode
	head := &main.ListNode{
		Val: val[0],
	}
	temp = head
	for i := 1; i < len(val); i++ {
		temp.Next = &main.ListNode{
			Val: val[i],
		}
		temp = temp.Next
	}
	return head
}

//func NewList(val ...int) *ListNode {
//	var head *ListNode
//    temp:=head.Next
//
//
//}
