package trees

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
 
 func maxDepth(root *TreeNode) int {
    if root == nil {
		return 0
	}

	l := maxDepth(root.Left)
	r :=  maxDepth(root.Right)
	if r > l {
		l = r
	}
	return l + 1
 }