package linklist
//编写一个程序，找到两个单链表相交的起始节点。
//ok
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	temp := headB
	for headA != nil {
		for temp != nil {
			if headA == temp {
				return headA
			}
			temp = temp.Next
		}
		temp = headB
		headA = headA.Next
	}
	return nil
}
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	nodeA := headA
	nodeB := headB
	for nodeA != nodeB {
		if nodeA != nil {
			nodeA = nodeA.Next
		} else {
			nodeA = headB //都走一边会交于交叉点
		}
		if nodeB != nil {
			nodeB = nodeB.Next
		} else {
			nodeB = headA
		}
	}
	return nodeA
}

