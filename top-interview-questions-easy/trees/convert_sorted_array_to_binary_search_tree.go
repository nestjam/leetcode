package trees

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func sortedArrayToBST(nums []int) *TreeNode {
// 	if len(nums) > 2 {
// 		i := len(nums) / 2
// 		return &TreeNode{
// 			Val: nums[i],
// 			Left: sortedArrayToBST(nums[0 : i]),
// 			Right: sortedArrayToBST(nums[i + 1 : ]),
// 		}
// 	} else if len(nums) == 2 {
// 		return &TreeNode{
// 			Val: nums[1],
// 			Left: &TreeNode {
// 				Val: nums[0],
// 			},
// 		}
// 	}
// 	return &TreeNode{Val: nums[0]}
// }

// Чуть более компактный вариант
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	i := len(nums) / 2
	return &TreeNode{
		Val:   nums[i],
		Left:  sortedArrayToBST(nums[0:i]),
		Right: sortedArrayToBST(nums[i+1:]),
	}
}
