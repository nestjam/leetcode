package spiralmatrixii

func generateMatrix(s int) [][]int {
	const (
		leftRight = 0
		topDown   = 1
		rightLeft = 2
		downTop   = 3
	)

	matrix := make([][]int, s)
	for i := 0; i < s; i++ {
		matrix[i] = make([]int, s)
	}

	n, m := s, s
	i, j := 0, 0
	k := 1
	direction := leftRight

	for n > 0 && m > 0 {
		switch direction {
		case leftRight:
			for l := 0; l < m; l++ {
				matrix[i][l+j] = k
				k++
			}
			i++
			j = j + m - 1
			n--
			direction = topDown
		case topDown:
			for l := 0; l < n; l++ {
				matrix[l+i][j] = k
				k++
			}
			i = i + n - 1
			j--
			m--
			direction = rightLeft
		case rightLeft:
			for l := 0; l < m; l++ {
				matrix[i][j-l] = k
				k++
			}
			i--
			j = j - m + 1
			n--
			direction = downTop
		case downTop:
			for l := 0; l < n; l++ {
				matrix[i-l][j] = k
				k++
			}
			i = i - n + 1
			j++
			m--
			direction = leftRight
		}
	}

	return matrix
}
