package Struct

type Stack interface {
	Pop()(rtn int, success bool)
	Push(a int) bool
	IsNull() bool
	IsFull() bool
	Show() []int
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
