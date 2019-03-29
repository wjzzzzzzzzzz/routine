package binTree

import (
	"testing"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}
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
	log.Println(tree)
	//tree.log()
	//depth := maxDepth(tree)
	//log.Println(depth)
}
func Test_hasPathSum(t *testing.T) {
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
	i := []int{}
	tree.traverse(&i)
	log.Println(i)
}
func Test_isValid(t *testing.T) {
	log.Println(isValid(10))
	NewBinTree(nil, 1, 1, nil)
}
func Test_NewTree(t *testing.T) {
	value := []int{1, 2, 3, 4, 5}
	TreeRoot := create(1, value)
	show(TreeRoot)
}
func Test_traverse(t *testing.T) {

}
func Test_levelOrderBottom(t *testing.T) {
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
	bottom := levelOrderBottom(tree)
	log.Println(bottom)
}
func Test_zigzagLevelOrder(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	bottom := zigzagLevelOrder(tree)
	log.Println(bottom)
}
