package cellswithoddvaluesinamatrix

func oddCells(m int, n int, indices [][]int) int {
	mi := make([]int, m)
	nj := make([]int, n)

	for ind := 0; ind < len(indices); ind++ {
		mi[indices[ind][0]]++
		nj[indices[ind][1]]++
	}

	odds := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (mi[i]+nj[j])%2 == 1 {
				odds++
			}
		}
	}

	return odds
}
