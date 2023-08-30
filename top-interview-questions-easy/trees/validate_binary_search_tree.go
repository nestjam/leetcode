package trees

import "fmt"

type stack struct {
	s []*TreeNode
}

func (c *stack) push(value *TreeNode) {
	c.s = append(c.s, value)
}

func (c *stack) pop() (value *TreeNode, err error) {
	len := len(c.s)
	if len > 0 {
		value = c.s[len-1]
		c.s = c.s[:len-1]
		return value, nil
	}
	err = fmt.Errorf("pop error: stack is empty")
	return
}

func (c *stack) empty() bool {
	return len(c.s) == 0
}

func isValidBST(root *TreeNode) bool {
	stack := new(stack)
	var pre *TreeNode
	for !stack.empty() || root != nil {
		for root != nil {
			stack.push(root)
			root = root.Left
		}
		root, _ = stack.pop()
		if pre != nil && root.Val <= pre.Val {
			return false
		}
		pre = root
		root = root.Right
	} 
	return true
}

// Выполнить обход дерева in-order, заполнить массив. Если массив упорядочен по возрастанию - бинарное дерево корректно 
/*
func isValidBST(root *TreeNode) bool {
	s := []int{}
	traverse(root, &s)
	for i := 1; i < len(s); i++ {
		if s[i-1] >= s[i] {
			return false
		}
	}
	return true
}

func traverse(node *TreeNode, s *[]int) {
	if node == nil {
		return
	}

	if node.Left != nil {
		traverse(node.Left, s)
	}

	*s = append(*s, node.Val)

	if node.Right != nil {
		traverse(node.Right, s)
	}
}
*/