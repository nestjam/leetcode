package squaresofasortedarray

func sortedSquares(n []int) []int {
	m := make([]int, 0, len(n))
	j := findFirstPositive(n)
	i := j - 1

	for i >= 0 || j < len(n) {
		if i < 0 {
			m = append(m, n[j]*n[j])
			j++
			continue
		}

		if j == len(n) {
			m = append(m, n[i]*n[i])
			i--
			continue
		}

		if -n[i] < n[j] {
			m = append(m, n[i]*n[i])
			i--
		} else {
			m = append(m, n[j]*n[j])
			j++
		}
	}

	return m
}

func findFirstPositive(n []int) int {
	for i := 0; i < len(n); i++ {
		if n[i] > 0 {
			return i
		}
	}
	return len(n)
}
