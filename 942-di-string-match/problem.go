package distringmatch

func diStringMatch(s string) []int {
	n := make([]int, len(s)+1)
	for i := 0; i < len(n); i++ {
		n[i] = i
	}

	for i := 0; i < len(n); i++ {
		cn := make([]int, len(n))
		copy(cn, n)
		cn[i] = -1
		if r, ok := next([]int{n[i]}, cn, 0, s); ok {
			return r
		}
	}

	return nil
}

func next(m []int, n []int, ind int, s string) ([]int, bool) {
	if ind == len(s) {
		return m, true
	}

	for i := 0; i < len(n); i++ {
		if n[i] == -1 {
			continue
		}

		if s[ind] == 'D' {
			if m[len(m)-1] < n[i] {
				continue
			}
		} else {
			if m[len(m)-1] > n[i] {
				continue
			}
		}

		cm := make([]int, len(m)+1)
		copy(cm, m)
		cm[len(m)] = n[i]
		cn := make([]int, len(n))
		copy(cn, n)
		cn[i] = -1
		if r, ok := next(cm, cn, ind+1, s); ok {
			return r, true
		}
	}

	return nil, false
}
