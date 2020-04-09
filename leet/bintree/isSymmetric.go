package binTree

/*给定一个二叉树，检查它是否是镜像对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

1
/ \
2   2
/ \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3
*/

//判断是否对称 =判断左子树和右子树是否对称 ，即 root
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSameTree(root.Right, invertTree(root.Left))
}
func test() {

}
