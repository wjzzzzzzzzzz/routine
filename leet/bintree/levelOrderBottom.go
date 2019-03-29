package binTree

//107
//给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var s []*TreeNode
	//新建的指针为空
	s = append(s, root.Left, root.Right)
	res := append([][]int{}, []int{root.Val})
	deals(s, &res)
	return res
}

func deals(s []*TreeNode, res *[][]int) {
	if len(s) == 0 {
		return
	}
	var r []int
	var m []*TreeNode
	for _, value := range s {
		if value != nil {
			r = append(r, value.Val)
			m = append(m, value.Left, value.Right)
		}
	}
	if len(r) != 0 {
		*res = append(*res, r)
	}
	deals(m, res)
}

