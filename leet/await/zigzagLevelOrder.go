package await

import "log"

//给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var s []*TreeNode
	//新建的指针为空
	s = append(s, root.Left, root.Right)
	res := append([][]int{}, []int{root.Val})
	deals1(s, &res, false)
	return res
}

func deals1(s []*TreeNode, res *[][]int, flag bool) {
	if len(s) == 0 {
		return
	}
	var r []int
	var m []*TreeNode
	if flag {
		for _, value := range s {
			if value != nil {
				r = append(r, value.Val)
				m = append(m, value.Left, value.Right)
			}
		}
	} else {
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] != nil {
				r = append(r, s[i].Val)
				m = append(m, s[i].Right, s[i].Left)
			}
		}
	}
	if len(r) != 0 {
		*res = append(*res, r)
	}
	log.Println(m,res,!flag)
	deals1(m, res,!flag)
}
