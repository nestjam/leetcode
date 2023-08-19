package binarytreeinordertraversal

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func inorderTraversal(root *TreeNode) []int {
// 	traversed := []int {}
// 	current := root
// 	stack := []*TreeNode { }
	
// 	for current != nil || len(stack) > 0 {
// 		for current != nil {
// 			stack = append(stack, current)
// 			current = current.Left
// 		}

// 		current, stack = stack[len(stack)-1], stack[:len(stack)-1]
// 		traversed = append(traversed, current.Val)
// 		current = current.Right
// 	}
// 	return traversed
// }


func inorderTraversal(root *TreeNode) []int {
	if (root == nil){
		return []int{}
	}

	traversed := append(inorderTraversal(root.Left), root.Val)
	return append(traversed, inorderTraversal(root.Right)...)
}

