package trees

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	return traverse([]*TreeNode{root}, [][]int{{root.Val}})
}

func traverse(level []*TreeNode, values [][]int) [][]int {
	var nextLevel []*TreeNode
	for _, node := range level {
		if node.Left != nil {
			nextLevel = append(nextLevel, node.Left)
		}
		if node.Right != nil {
			nextLevel = append(nextLevel, node.Right)
		}
	}
	if len(nextLevel) > 0 {
		nextValues := make([]int, len(nextLevel))
		for i, node := range nextLevel {
			nextValues[i] = node.Val
		}
		values = append(values, nextValues)
		return traverse(nextLevel, values)
	}
	return values
}
