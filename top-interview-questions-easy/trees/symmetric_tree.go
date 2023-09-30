package trees

//  type TreeNode struct {
// 	Val int
// 	Left *TreeNode
// 	Right *TreeNode
//  }

func isSymmetric(root *TreeNode) bool {
	return compare(root.Left, root.Right)
 }

func compare(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	} else {
		return left.Val == right.Val && compare(left.Left, right.Right) && compare(left.Right, right.Left)
	}
}