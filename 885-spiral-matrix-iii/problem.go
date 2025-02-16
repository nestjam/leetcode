package spiralmatrixiii

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	m := make([][]int, rows*cols)

	iterator := newSpiralIterator(rStart, cStart)
	i := 0
	m[i] = iterator.current()
	i++

	for i < len(m) {
		iterator.moveNext()
		cur := iterator.current()

		if cur[0] < 0 || cur[0] >= rows || cur[1] < 0 || cur[1] >= cols {
			continue
		}

		m[i] = cur
		i++
	}

	return m
}

type spiralIterator struct {
	direction byte
	i         int
	j         int
	step      int
	steps     int
	round     int
	dir       [][]int
}

func newSpiralIterator(i, j int) spiralIterator {
	return spiralIterator{
		dir:   [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}},
		i:     i,
		j:     j,
		steps: 1,
	}
}

func (s *spiralIterator) current() []int {
	return []int{s.i, s.j}
}

func (s *spiralIterator) moveNext() {
	s.i += s.dir[s.direction][0]
	s.j += s.dir[s.direction][1]

	s.step++

	if s.step == s.steps {
		s.step = 0
		s.round++
		s.direction = (s.direction + 1) % 4

		if s.round == 2 {
			s.steps++
			s.round = 0
		}
	}
}
