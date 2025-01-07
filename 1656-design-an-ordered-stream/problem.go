package designanorderedstream

type OrderedStream struct {
	stream []string
	p      int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		stream: make([]string, n),
		p:      0,
	}
}

func (s *OrderedStream) Insert(k int, v string) []string {
	k--

	prev := s.p
	s.stream[k] = v

	if s.p == k {
		for s.p < len(s.stream) {
			if s.stream[s.p] == "" {
				break
			}

			s.p++
		}
	}
	return s.stream[prev:s.p]
}
