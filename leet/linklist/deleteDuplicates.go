package linklist

//给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

func deleteDuplicates(head *ListNode) *ListNode {
	temp := head
	for temp != nil {
		if temp.Next==nil{
			break
		}
		if temp.Val==temp.Next.Val{
			temp.Next=temp.Next.Next
		}else{
			temp=temp.Next
		}
	}
	return head
}

