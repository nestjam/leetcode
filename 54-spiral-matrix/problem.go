package spiralmatrix

func spiralOrder(matrix [][]int) []int {
	n := len(matrix)
	m := len(matrix[0])
	line := make([]int, n*m)

	readLeftRight(matrix, 0, 0, n, m, line, 0)

	return line
}

func readLeftRight(matrix [][]int, i, j, n, m int, line []int, k int) {
	if n == 0 || m == 0 {
		return
	}

	for l := 0; l < m; l++ {
		line[k] = matrix[i][l+j]
		k++
	}

	readTopDown(matrix, i+1, j+m-1, n-1, m, line, k)
}

func readTopDown(matrix [][]int, i, j, n, m int, line []int, k int) {
	if n == 0 || m == 0 {
		return
	}
	
	for l := 0; l < n; l++ {
		line[k] = matrix[l+i][j]
		k++
	}

	readRightLeft(matrix, i+n-1, j-1, n, m-1, line, k)
}

func readRightLeft(matrix [][]int, i, j, n, m int, line []int, k int) {
	if n == 0 || m == 0 {
		return
	}
	
	for l := 0; l < m; l++ {
		line[k] = matrix[i][j-l]
		k++
	}

	readDownTop(matrix, i-1, j-m+1, n-1, m, line, k)
}

func readDownTop(matrix [][]int, i, j, n, m int, line []int, k int) {
	if n == 0 || m == 0 {
		return
	}
	
	for l := 0; l < n; l++ {
		line[k] = matrix[i-l][j]
		k++
	}

	readLeftRight(matrix, i-n+1, j+1, n, m-1, line, k)
}
