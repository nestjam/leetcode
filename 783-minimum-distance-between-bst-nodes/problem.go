package minimumdistancebetweenbstnodes

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type data struct {
	minDiff int
	prevNode *TreeNode
}

func minDiffInBST(node *TreeNode) int {
	data := &data{math.MaxInt32, nil}
	traverseBST(node, data)
	return data.minDiff
}

func traverseBST(node *TreeNode, data *data) {
	if node.Left != nil {
		traverseBST(node.Left, data)
	}

	if data.prevNode != nil {
		diff := node.Val - data.prevNode.Val
		if data.minDiff > diff {
			data.minDiff = diff
		}
	}
	data.prevNode = node

	if node.Right != nil {
		traverseBST(node.Right, data)
	}
}
