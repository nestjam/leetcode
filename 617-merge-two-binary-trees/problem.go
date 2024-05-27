package mergetwobinarytrees

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root2 == nil {
		return root1
	}

	if root1 == nil {
		return root2
	}

	root1.Val += root2.Val

	if root1.Left == nil {
		root1.Left = root2.Left
	} else if root2.Left != nil {
		root1.Left = mergeTrees(root1.Left, root2.Left)
	}

	if root1.Right == nil {
		root1.Right = root2.Right
	} else if root2.Right != nil {
		root1.Right = mergeTrees(root1.Right, root2.Right)
	}

	return root1
}
