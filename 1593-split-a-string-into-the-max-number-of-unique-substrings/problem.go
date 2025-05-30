package splitastringintothemaxnumberofuniquesubstrings

var m int

func maxUniqueSplit(s string) int {
	m = 1

	for i := 1; i < len(s); i++ {
		c := newHead(i)
		split(&c, s)
	}

	return m
}

func split(p *substring, s string) {
	if p.end == len(s) {
		m = max(m, p.len())
		return
	}

	c := newSubstring(p.end, p)

	for i := p.end; i < len(s); i++ {
		c.end = i + 1
		if !p.contains(c, s) {
			split(&c, s)
		}
	}
}

type substring struct {
	prev *substring
	end  int
}

func newSubstring(end int, prev *substring) substring {
	return substring{end: end, prev: prev}
}

func newHead(end int) substring {
	return newSubstring(end, nil)
}

func (s *substring) value(t string) string {
	start := 0
	if s.prev != nil {
		start = s.prev.end
	}
	return t[start:s.end]
}

func (s *substring) contains(c substring, t string) bool {
	for s != nil {
		if c.value(t) == s.value(t) {
			return true
		}

		s = s.prev
	}

	return false
}

func (s *substring) len() int {
	n := 0
	for s != nil {
		n++
		s = s.prev
	}
	return n
}
