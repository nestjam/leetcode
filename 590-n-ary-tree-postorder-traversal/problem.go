package narytreepostordertraversal

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	if root == nil {
		return []int{}
	}

	const capacity = 5000 // max count is 10000
	values := make([]int, 0, capacity)
	postorderRecursively(root, &values)
	return values
}

func postorderRecursively(root *Node, values *[]int) {
	for i := 0; i < len(root.Children); i++ {
		postorderRecursively(root.Children[i], values)
	}
	*values = append(*values, root.Val)
}
