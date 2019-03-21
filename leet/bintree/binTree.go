package binTree

import (
	"math"
	"log"
)

/*public void mirror(TreeNode root){
//二叉树的镜像翻转
if(root == null){
return ;
}
if(root.left == null && root.right == null){
return ;
}
TreeNode temp = null;
temp = root.left;
root.left = root.right;
root.right = temp;

if(root.left !=null){
mirror(root.left);
}

if(root.right != null){
mirror(root.right);
}
}
*/

func Mirror(root *TreeNode) {
	if root == nil || root.Left == nil || root.Right == nil {
		return
	}
	temp := root
	temp.Left, temp.Right = temp.Right, temp.Left
	if temp.Left != nil {
		Mirror(temp.Left)
	}
	if temp.Right != nil {
		Mirror(temp.Right)
	}
}

func (t *TreeNode) log() {
	s := make([]int, 0, 500)
	//var s []int
	t.mergeValue(s)
	log.Println(s)
}
func (t *TreeNode) mergeValue(a []int) {
	if t != nil {
		log.Println("之前", a)
		a = append(a, t.Val)
		log.Printf("%p", a)
		log.Println("之后", a)
		log.Println(t.Val)
		t.Left.mergeValue(a)
		t.Right.mergeValue(a)
	}
}
func NewBinTree(nodes ...interface{}) *TreeNode {
	if !isValid(len(nodes)) {
		return nil
	}
	//root:=&TreeNode{
	//}

	for i := 1; i < len(nodes); i++ {
		switch nodes[i].(type) {
		case int:
		case nil:
		}
	}
	return nil
}

type Node int

func isValid(l int) bool {
	x := int(math.Sqrt(float64(l))) + 1
	if 2*l == (x)*(x+1) {
		return true
	}
	return false
}

