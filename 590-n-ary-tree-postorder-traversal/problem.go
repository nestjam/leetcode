package narytreepostordertraversal

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	if root == nil {
		return []int{}
	}

	return postorderRecursively(root)
}

func postorderRecursively(root *Node) []int {
	values := make([]int, 0)

	for i := 0; i < len(root.Children); i++ {
		values = append(values, postorderRecursively(root.Children[i])...)
	}

	return append(values, root.Val)
}
