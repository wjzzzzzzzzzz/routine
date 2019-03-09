package linklist

import (
	"testing"
	"log"
)

func Test_addTwoNumbers(t *testing.T) {
	node1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 5,
			},
		},
	}
	node2 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 6,
			},
		},
	}
	lists := mergeTwoLists(node1, node2)
	for {
		log.Println(lists.Val)
		if lists.Next != nil {
			lists = lists.Next

		} else {
			break
		}

	}
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
