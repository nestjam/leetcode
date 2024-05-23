package binarytreepreordertraversal

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//recursive
// func preorderTraversal(root *TreeNode) []int {
// 	if root == nil {
// 		return []int{}
// 	}

// 	traversed := append([]int{root.Val}, preorderTraversal(root.Left)...)
// 	return append(traversed, preorderTraversal(root.Right)...)
// }

func preorderTraversal(root *TreeNode) []int {
	traversed := []int{}
	current := root
	stack := []*TreeNode{}

	for current != nil || len(stack) > 0 {
		for current != nil {
			traversed = append(traversed, current.Val)
			stack = append(stack, current)
			current = current.Left
		}

		current, stack = stack[len(stack)-1], stack[:len(stack)-1]
		current = current.Right
	}
	return traversed
}
