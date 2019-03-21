package linklist

/*给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807*/

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode
	var temp *ListNode //定义了一个指针但是没有申请内存 ，导致判断的时候出问题 ，先判断
	temp = res
	var x, y, sum, carry int
	p, q := l1, l2
	for p != nil || q != nil { //只要一条可以
		if p == nil {
			x = 0
		} else {
			x = p.Val
		}
		if q == nil {
			y = 0
		} else {
			y = q.Val
		}
		sum = x + y + carry
		carry = sum / 10
		if res == nil {
			res = &ListNode{
			}
			temp = &ListNode{Val: sum % 10,}
			res.Next = temp
		} else {
			temp.Next = &ListNode{Val: sum % 10}
			temp = temp.Next
		}
		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}
	if carry == 1 {
		temp.Next = &ListNode{
			Val: 1,
		}
	}
	return res.Next
}
