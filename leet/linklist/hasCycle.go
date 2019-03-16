package linklist

//给定一个链表，判断链表中是否有环。
//
//为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
//ok

//快慢指针
func hasCycle(head *ListNode) bool {
	slow := head
	fast := head
	for slow!=nil &&fast!=nil{
		slow=slow.Next
		if fast.Next==nil{
			return false
		}
		fast=fast.Next.Next
		if slow==fast{
			return true
		}
	}
	return false
}
