package linklist

//将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}
	var result *ListNode
	for l1 == nil && l2 == nil {
		if l1.Val <= l1.Val {
			result.Next = l1
			l1 = l1.Next
			result = result.Next
		} else {
			result.Next = l2
			l2 = l2.Next
			result = result.Next
		}
	}
	if l1 == nil {
		result.Next = l2
	}
	if l2 == nil {
		result.Next = l1
	}
	return result
}
