package binTree

import (
	"math"
	"log"
	"fmt"
)

func Mirror(root *TreeNode) {

	invertTree(root)

}


func NewBinTree(nodes ...interface{}) *TreeNode {
	if !isValid(len(nodes)) {
		return nil
	}
	for i := 1; i < len(nodes); i++ {
		switch nodes[i].(type) {
		case int:

		case nil:
		}
	}
	return nil
}

func create(index int, value []int) (T *TreeNode) {
	T = &TreeNode{}
	T.Val = value[index-1]
	fmt.Printf("index %v \n", index)
	if index < len(value)-1 && 2*index <= len(value) && 2*index+1 <= len(value) {
		T.Left = create(2*index, value)
		T.Right = create(2*index+1, value)
	}
	return T
}

func show(treeNode *TreeNode) {
	if treeNode != nil {
		fmt.Printf("%v ", treeNode.Val)
		if treeNode.Left != nil {
			show(treeNode.Left)
		}
		if treeNode.Right != nil {
			show(treeNode.Right)
		}
	} else {
		return
	}
}
func (t *TreeNode) traverse(s *[]int) {
	if t == nil {
		return
	}
	*s= append(*s, t.Val)
	log.Printf("%p         %v,%v",s,s,t.Val)
	t.Left.traverse(s)
	t.Right.traverse(s)
}
type Node int
func isValid(l int) bool {
	x := int(math.Sqrt(float64(l))) + 1
	if 2*l == (x)*(x+1) {
		return true
	}
	return false
}

//************************************************************
////二叉搜索树
//type BinarySearchTree struct {
//	Root *TreeNode
//}
