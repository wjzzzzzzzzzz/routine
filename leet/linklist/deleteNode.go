package linklist

//请编写一个函数，使其可以删除某个链表中给定的（非末尾）节点，你将只被给定要求被删除的节点。

func deleteNode(node *ListNode) {
		node.Val=node.Next.Val
		node.Next=node.Next.Next
}
