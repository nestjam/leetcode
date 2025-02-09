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

type direction = byte

const (
	leftRight direction = iota
	topDown
	rightLeft
	downTop
)

type spiralIterator struct {
	direction direction
	i         int
	j         int
	side      int
	round     int
}

func newSpiralIterator(i, j int) spiralIterator {
	return spiralIterator{
		direction: leftRight,
		i:         i,
		j:         j,
		side:      1,
		round:     2,
	}
}

func (s *spiralIterator) current() []int {
	return []int{s.i, s.j}
}

func (s *spiralIterator) moveNext() {
	s.round--

	switch s.direction {
	case leftRight:
		s.j++

		if s.round == s.side {
			s.direction = topDown
		}
	case topDown:
		s.i++

		if s.round == 0 {
			s.direction = rightLeft
			s.side++
			s.round = 2 * s.side
		}
	case rightLeft:
		s.j--

		if s.round == s.side {
			s.direction = downTop
		}
	case downTop:
		s.i--

		if s.round == 0 {
			s.direction = leftRight
			s.side++
			s.round = 2 * s.side
		}
	}
}
