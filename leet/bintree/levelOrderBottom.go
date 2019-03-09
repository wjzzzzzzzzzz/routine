package binTree

//给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
//func levelOrderBottom(root *TreeNode) [][]int {
//
//	temp:=root
//	for temp!=nil{
//		m:=append(getLR(temp.Left),getLR(temp.Right)...)
//
//	}
//
//
//
//}
//
//
//func getLR(root *TreeNode) []int {
//	var result []int
//	if root.Left != nil {
//		result = append(result, root.Left.Val)
//	}
//	if root.Right != nil {
//		result = append(result, root.Right.Val)
//	}
//	return result
//
//}
//
//func levelOrderBottom(root *TreeNode) [][]int{
//	if root==nil{
//		return nil
//	}
//	temp:=root
//
//	for temp!=nil{
//		append()getValue(temp),getValue(temp.Left),getValue(temp.Right)
//
//	}
//
//}
//
//
//func getValue(root *TreeNode)[]int{
//	return append(make([]int,0),root.Left.Val,root.Right.Val)
//}