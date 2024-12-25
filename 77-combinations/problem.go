package combinations

func combine(n int, k int) [][]int {
	c := make([][]int, 0)

	for i := 1; i <= n; i++ {
		if s, ok := combineRecursively(i, n, k); ok {
			c = append(c, s...)
		}
	}

	return c
}

func combineRecursively(m, n, k int) ([][]int, bool) {
	if n-m+1 < 1 {
		return nil, false
	}

	if k == 1 {
		return [][]int{{m}}, true
	} else {
		c := make([][]int, 0)

		for i := m + 1; i <= n; i++ {
			if s, ok := combineRecursively(i, n, k-1); ok {
				c = append(c, s...)
			}
		}

		for i := 0; i < len(c); i++ {
			c[i] = append(c[i], m)
		}

		return c, len(c) > 0
	}
}
