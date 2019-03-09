package binTree

//给定一个二叉树，找出其最大深度。
//
//二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
//
//说明: 叶子节点是指没有子节点的节点。

//层序遍历  父节点出 子节点进入
//func maxDepth1(root *TreeNode) int {
//
//
//}
//ok
func maxDepth(root *TreeNode) int {
	if root==nil{
		return 0
	}
	left:= maxDepth(root.Left)
	right:= maxDepth(root.Right)
	if left>right{
		return  left+1
	}else{
		return right+1
	}

}
