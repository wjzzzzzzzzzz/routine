package binTree
//翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil{
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)//翻转过的右边和左边
	return root
}
