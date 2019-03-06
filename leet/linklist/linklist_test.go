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
	lists := mergeTwoLists(node1,node2)
	for {
		log.Println(lists.Val)
		if lists.Next!=nil{
			lists=lists.Next

		}else{
			break
		}

	}
}