package binTree

import (
	"testing"
	"log"
)

func Test_maxDepth(t *testing.T) {
	tree:=&TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 5,
				},
				Left:&TreeNode{
					Val: 6,
				},
			},
			Left:&TreeNode{
				Val: 7,
			},
		},
		Left: &TreeNode{
			Val: 2,
		},
	}
	depth := maxDepth(tree)
	log.Println(depth)

}
func Test_hasPathSum(t *testing.T) {
	tree:=&TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 100,
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 5,
				},
				Left:&TreeNode{
					Val: 6,
				},
			},
			Left:&TreeNode{
				Val: 200,
			},
		},
		Left: &TreeNode{
			Val: 100,
		},
	}
	depth := hasPathSum(tree,101)
	log.Println(depth)

}