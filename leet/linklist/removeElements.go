package linklist

//删除链表中等于给定值 val 的所有节点。
//
//示例:
//
//输入: 1->2->6->3->4->5->6, val = 6
//输出: 1->2->3->4->5
func removeElements(head *ListNode, val int) *ListNode {
	if head==nil{
		return nil
	}
	temp := head
	for temp.Next != nil {//第一个节点先不管
		if temp.Next.Val==val{//如果下一个节点等于
			temp.Next=temp.Next.Next//那么跳过这个节点
		}else {
			temp = temp.Next
		}
	}
	if head.Val==val{
		head=head.Next
	}
	return head
}
func removeElements1(head *ListNode, val int) *ListNode {
	temp := head
	for temp.Next!= nil {
		if temp.Next.Val==val{
			temp.Next=temp.Next.Next

		}else {
			temp = temp.Next

		}
	}
	return head
}
