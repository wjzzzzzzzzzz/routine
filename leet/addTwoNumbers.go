package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
//如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
//您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//示例：
//
//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	temp := l1
	var m []int
	for {
		m = append(m, temp.Val)
		if temp.Next != nil {
			temp = temp.Next

		} else {
			break
		}
	}
	l := len(m)
	ints := make([]int, l)
	for i := len(m) - 1; i >= 0; i-- {
		ints = append(ints, m[i])
	}
	return nil
}

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	temp1 := l1
//	temp2 := l2
//	var l1value, l2value int
//	for ; temp1.Next != nil && temp2.Next != nil; {
//		l1value += temp1.Val
//		l2value += temp2.Val
//		temp1 = temp1.Next
//		temp2 = temp2.Next
//		temp1.Val = temp1.Val * 10
//		temp2.Val = temp2.Val * 10
//	}
//	if temp1.Next != nil{
//
//	}
//
//
//}
