package binTree

import (
	"testing"
	"log"
)

func Test_maxDepth(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 5,
				},
				Left: &TreeNode{
					Val: 6,
				},
			},
			Left: &TreeNode{
				Val: 7,
			},
		},
		Left: &TreeNode{
			Val: 2,
		},
	}
	tree.log()
	//depth := maxDepth(tree)
	//log.Println(depth)
}
func Test_hasPathSum(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 100,
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 5,
				},
				Left: &TreeNode{
					Val: 6,
				},
			},
			Left: &TreeNode{
				Val: 200,
			},
		},
		Left: &TreeNode{
			Val: 100,
		},
	}
	depth := hasPathSum(tree, 101)
	log.Println(depth)
}
func Test_isValid(t *testing.T) {
	log.Println(isValid(10))
	NewBinTree(nil,1,1,nil)
}
func Test_NewTree(t *testing.T) {
	value := []int{1, 2, 3, 4, 5}
	TreeRoot := create(1, value)
	show(TreeRoot)
}