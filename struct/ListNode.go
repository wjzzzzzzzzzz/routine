package Struct

//

func NewList(val ...int) *ListNode {
	var temp *ListNode
	head := &ListNode{
		Val: val[0],
	}
	temp = head
	for i := 1; i < len(val); i++ {
		temp.Next = &ListNode{
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
