package narytreepreordertraversal

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}

	const capacity = 5000 // max count is 10000
	values := make([]int, 0, capacity)
	preorderRecursively(root, &values)
	return values
}

func preorderRecursively(root *Node, values *[]int) {
	*values = append(*values, root.Val)
	for i := 0; i < len(root.Children); i++ {
		preorderRecursively(root.Children[i], values)
	}
}
