package spiralmatrix

func spiralOrder(matrix [][]int) []int {
	const (
		leftRight = 0
		topDown   = 1
		rightLeft = 2
		downTop   = 3
	)

	n := len(matrix)
	m := len(matrix[0])
	line := make([]int, n*m)

	i, j, k := 0, 0, 0
	direction := leftRight

	for n > 0 && m > 0 {
		switch direction {
		case leftRight:
			for l := 0; l < m; l++ {
				line[k] = matrix[i][l+j]
				k++
			}
			i++
			j = j + m - 1
			n--
			direction = topDown
		case topDown:
			for l := 0; l < n; l++ {
				line[k] = matrix[l+i][j]
				k++
			}
			i = i + n - 1
			j--
			m--
			direction = rightLeft
		case rightLeft:
			for l := 0; l < m; l++ {
				line[k] = matrix[i][j-l]
				k++
			}
			i--
			j = j - m + 1
			n--
			direction = downTop
		case downTop:
			for l := 0; l < n; l++ {
				line[k] = matrix[i-l][j]
				k++
			}
			i = i - n + 1
			j++
			m--
			direction = leftRight
		}
	}

	return line
}
