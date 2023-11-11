package design

type node struct {
	val  int
	min  int
	next *node
}

type MinStack struct {
	top *node
}

func NewMinStack() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(val int) {
	min := val
	if s.top != nil && min > s.top.min {
		min = s.top.min
	}
	node := &node{val: val, min: min}
	node.next = s.top
	s.top = node
}

func (s *MinStack) Pop() {
	s.top = s.top.next
}

func (s *MinStack) Top() int {
	return s.top.val
}

func (s *MinStack) GetMin() int {
	return s.top.min
}
