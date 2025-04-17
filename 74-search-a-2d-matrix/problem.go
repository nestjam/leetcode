package searcha2dmatrix

func searchMatrix(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])

	for i := 0; i < n; i++ {
		if matrix[i][0] > target ||
			matrix[i][m-1] < target {
			continue
		}

		for j := 0; j < m; j++ {
			if matrix[i][j] == target {
				return true
			}
		}

		break
	}

	return false
}
