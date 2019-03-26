package linklist

/*请判断一个链表是否为回文链表。*/

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast := head
	slow := head
	temp := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next //fast==nil
		slow = slow.Next
	}
	list := reverseList(slow)
	for list != temp &&temp!=nil {
		if list.Val != temp.Val {
			return false
		}
		list = list.Next
		temp = temp.Next
	}
	return true
}
