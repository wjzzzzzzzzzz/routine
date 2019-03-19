package linklist

import "log"

func NewList(isHead bool, val ...int) *ListNode {
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
	if isHead {
		return head
	} else {
		return temp
	}

}
func (l *ListNode) log() {
	if l==nil{
		return
	}
	var m []int
	for {
		m = append(m, l.Val)
		if l.Next != nil {
			l = l.Next
		} else {
			break
		}

	}
	log.Println(m)
}
