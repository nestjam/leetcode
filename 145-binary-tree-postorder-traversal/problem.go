package binarytreepostordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// recursive
// func postorderTraversal(root *TreeNode) []int {
// 	if (root == nil){
// 		return []int {}
// 	}

// 	traversed := append(postorderTraversal(root.Left), postorderTraversal(root.Right)...)
// 	return append(traversed, root.Val)
// }

// iterative
func postorderTraversal(root *TreeNode) []int {
	traversed := []int{}
	current := root
	stack := []*TreeNode{}

	for current != nil || len(stack) > 0 {
		if current != nil {
			traversed = append([]int{current.Val}, traversed...) //append(traversed, current.Val)
			stack = append(stack, current)
			current = current.Right
		} else {
			current, stack = stack[len(stack)-1], stack[:len(stack)-1]
			current = current.Left
		}
	}
	return traversed
}
