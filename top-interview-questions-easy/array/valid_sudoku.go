package array

func isValidSudoku(board [][]byte) bool {
	const dot byte = '.'
	const one = '1'
	const size = 9
	counters := make([]byte, size)

	// by rows
	for i := 0; i < size; i++ {
		clear(counters)
		for j := 0; j < size; j++ {
			if board[i][j] != dot {
				counters[board[i][j]-one]++
			}
		}
		if check(counters) == false {
			return false
		}
	}

	// by columns
	for i := 0; i < size; i++ {
		clear(counters)
		for j := 0; j < size; j++ {
			if board[j][i] != dot {
				counters[board[j][i]-one]++
			}
		}
		if check(counters) == false {
			return false
		}
	}

	// by 3x3 squares
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			clear(counters)

			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					cell := board[3*i+k][3*j+l]
					if cell != dot {
						counters[cell-one]++
					}
				}
			}

			if check(counters) == false {
				return false
			}
		}
	}

	return true
}

func clear(s []byte) {
	for i := 0; i < len(s); i++ {
		s[i] = 0
	}
}

func check(s []byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > 1 {
			return false
		}
	}
	return true
}
