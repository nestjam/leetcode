package distringmatch

type node struct {
	val  int
	prev *node
}

func diStringMatch(s string) []int {
	for i := 0; i < len(s)+1; i++ {
		if n, ok := next(&node{val: i}, 0, s); ok {
			r := make([]int, len(s)+1)
			for j := len(r) - 1; j >= 0; j-- {
				r[j] = n.val
				n = n.prev
			}
			return r
		}
	}

	return nil
}

func next(n *node, ind int, s string) (*node, bool) {
	if ind == len(s) {
		return n, true
	}

	for i := 0; i < len(s)+1; i++ {
		if contains(n, i) {
			continue
		}

		if s[ind] == 'D' {
			if n.val < i {
				continue
			}
		} else {
			if n.val > i {
				continue
			}
		}

		m := &node{val: i, prev: n}
		if m, ok := next(m, ind+1, s); ok {
			return m, true
		}
	}

	return nil, false
}

func contains(n *node, v int) bool {
	for n != nil {
		if n.val == v {
			return true
		}
		n = n.prev
	}
	return false
}
