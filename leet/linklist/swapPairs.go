package linklist
//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//
//你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	other := &ListNode{
		Next: head,
	}
	t1 := head
	t2 := t1.Next
	t1.Next, t2.Next, other.Next = t2.Next, t1, t2
	t1.Next = swapPairs(t1.Next)
	return other.Next
}
