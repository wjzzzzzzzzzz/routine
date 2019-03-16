package await

//给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
import (

	"routine/leet/bintree"
)

func levelOrderBottom(root *binTree.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	temp := root

	for temp != nil {
		getValue(temp)



	}

}

func getValue(root *binTree.TreeNode) []int {
	return append(make([]int, 0), root.Left.Val, root.Right.Val)
}
