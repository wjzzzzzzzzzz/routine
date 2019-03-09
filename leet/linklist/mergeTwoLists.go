package linklist

//将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil{
		return l2
	}
	if l2 == nil{
		return l1
	}
	var res *ListNode

	//返回下个节点
	if l1.Val >= l2.Val{
		res = l2
		res.Next = mergeTwoLists(l1,l2.Next)
	}else{
		res = l1
		res.Next = mergeTwoLists(l1.Next,l2)
	}
	return res
}