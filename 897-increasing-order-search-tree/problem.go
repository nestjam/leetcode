package increasingordersearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	node := root

	if root.Left != nil {
		node = increasingBST(root.Left)
		root.Left = nil
		last := node
		for last.Right != nil {
			last = last.Right
		}
		last.Right = root
	}

	if root.Right != nil {
		root.Right = increasingBST(root.Right)
	}

	return node
}
