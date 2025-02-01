package squaresofasortedarray

func sortedSquares(n []int) []int {
	m := make([]int, len(n))
	i := 0
	j := len(n) - 1

	for k := len(n) - 1; k >= 0; k-- {
		if abs(n[i]) > n[j] {
			m[k] = n[i] * n[i]
			i++
		} else {
			m[k] = n[j] * n[j]
			j--
		}
	}

	return m
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
