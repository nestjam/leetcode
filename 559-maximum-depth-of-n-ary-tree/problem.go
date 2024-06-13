package maximumdepthofnarytree

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	n := 0
	for i := 0; i < len(root.Children); i++ {
		m := maxDepth(root.Children[i])
		if m > n {
			n = m
		}
	}

	return n + 1
}
