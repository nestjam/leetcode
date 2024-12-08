package validarrangementofpairs

func validArrangement(pairs [][]int) [][]int {
	inOutDegree := make(map[int]int)
	graph := make(map[int][]int)

	for i := 0; i < len(pairs); i++ {
		start := pairs[i][0]
		end := pairs[i][1]
		graph[start] = append(graph[start], end)
		inOutDegree[start]++
		inOutDegree[end]--
	}

	stack := stack{}
	circuit := make([]int, 0)

	cur_v := pairs[0][0]
	for node, degree := range inOutDegree {
		if degree == 1 {
			cur_v = node
			break
		}
	}

	stack.Push(cur_v)

	for !stack.IsEmpty() {
		if len(graph[cur_v]) > 0 {
			stack.Push(cur_v)
			next_v := graph[cur_v][len(graph[cur_v])-1]
			graph[cur_v] = graph[cur_v][:len(graph[cur_v])-1]
			cur_v = next_v
		} else {
			circuit = append(circuit, cur_v)
			cur_v = stack.Pop()
		}
	}

	j := 0
	path := make([][]int, len(circuit)-1)

	for i := len(circuit) - 1; i > 0; i-- {
		path[j] = []int{circuit[i], circuit[i-1]}
		j++
	}

	return path
}

type stack struct {
	items []int
}

func (s *stack) Push(data int) {
	s.items = append(s.items, data)
}

func (s *stack) Pop() int {
	if s.IsEmpty() {
		panic("stack is empty")
	}

	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top
}

func (s *stack) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}
