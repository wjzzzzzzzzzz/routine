package linklist

import (
	"testing"
	"log"
)

func NewList(isHead bool, val ...int) *ListNode {
	var temp *ListNode
	head := &ListNode{
		Val: val[0],
	}
	temp = head
	for i := 1; i < len(val); i++ {
		temp.Next = &ListNode{
			Val: val[i],
		}
		temp = temp.Next
	}
	if isHead {
		return head
	} else {
		return temp
	}

}
func (l *ListNode) log() {
	var m []int
	for {
		m=append(m,l.Val)
		if l.Next != nil {
			l = l.Next
		}else{
			break
		}

	}
	log.Println(m)
}

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
	list1 := NewList(true, 1,3,4)
	list := reverseList(list1)
	list.log()
}
