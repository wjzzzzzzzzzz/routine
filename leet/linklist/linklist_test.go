package linklist

import (
	"testing"
	"log"
)

func Test_addTwoNumbers(t *testing.T) {
	l1 := NewList(true, 1, 0, 0, 0, 0, 0, 8)
	l2 := NewList(true, 0)
	res := addTwoNumbers1(l1, l2)
	res.log()

}

func Test_deleteDuplicates(t *testing.T) {
	node1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 6,
							Next: &ListNode{
								Val: 7,
								Next: &ListNode{
									Val: 7,
								},
							},
						},
					},
				},
			},
		},
	}
	node := deleteDuplicates(node1)
	for node != nil {
		log.Println(node.Val)
		node = node.Next
	}
}

func Test_getIntersectionNode(t *testing.T) {
	list1 := NewList(true, 1)
	list2 := NewList(true, 1)
	//node := &ListNode{Val: 7}
	//
	//list1.ne
	//list2.Next = node
	//node.Next=&ListNode{Val: 8}
	log.Println(getIntersectionNode(list1, list2))

}
func Test_reverseList(t *testing.T) {
	list1 := NewList(true, 1, 3, 4)
	list := reverseList(list1)
	list.log()
}
func Test_removeElements(t *testing.T) {
	list1 := NewList(true, 4, 3, 3, 3, 4, 3)
	removeElements(list1, 3).log()
}
func Test_isPalindrome(t *testing.T) {

	//list1 := NewList(true, 1,2,3,2,1)
	list1 := NewList(true, 0, 1)
	log.Println(isPalindrome(list1))
}
func Test_deleteNode(t *testing.T) {
	list1 := NewList(true, 0, 1,2,3)
	deleteNode(list1.Next)
	list1.log()
}
