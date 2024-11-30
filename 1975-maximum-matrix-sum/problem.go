package maximummatrixsum

func maxMatrixSum(m [][]int) int64 {
	var sum int64
	n := 0
	a := m[0][0]
	if a < 0 {
		a = -a
	}

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			v := m[i][j]

			if v < 0 {
				n++
				v = -v
			}

			if a > v {
				a = v
			}

			sum += int64(v)
		}
	}

	if n%2 == 1 {
		sum -= 2 * int64(a)
	}

	return sum
}
