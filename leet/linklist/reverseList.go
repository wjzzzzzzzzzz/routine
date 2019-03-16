package linklist

//ok

//反转一个单链表。
//func reverseList(head *ListNode) *ListNode {
//	if head.Next == nil {
//		head.Next = head
//		return head
//	}
//
//	return reverseList(head.Next)
//}
//func reverseList(head *ListNode) *ListNode {
//
//	return rever(head, head.Next)
//
//}
//func rever(a, b *ListNode) (c,d *ListNode) {
//	if b.Next == nil {
//		b.Next = a
//		re
//	}
//
//	return rever(a.Next, b.Next)
//}

func reverseList(head *ListNode) *ListNode {
	var result *ListNode
	p := head
	for p != nil {
		result, p, p.Next = p, p.Next, result
	}
	return result
}

//func reverseList(head *ListNode) *ListNode {
//	var result *ListNode
//	p := head
//
//	for p.Next != nil {
//		p = p.Next
//	}
//
//
//	return result
//}
