package gameoflife

const (
	dead    = 0
	live    = 1
	dying   = 2
	nascent = 3
)

func gameOfLife(board [][]int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			updateState(board, i, j)
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			commitState(board, i, j)
		}
	}
}

func commitState(board [][]int, i, j int) {
	if board[i][j] == dying {
		board[i][j] = dead
	} else if board[i][j] == nascent {
		board[i][j] = live
	}
}

func updateState(board [][]int, i, j int) {
	liveNeigboors := countLiveNeighboors(board, i, j)

	if board[i][j] == live {
		if liveNeigboors < 2 {
			board[i][j] = dying
		} else if liveNeigboors == 2 || liveNeigboors == 3 {
			board[i][j] = live
		} else if liveNeigboors > 3 {
			board[i][j] = dying
		}
	} else if board[i][j] == dead {
		if liveNeigboors == 3 {
			board[i][j] = nascent
		}
	} else {
		panic("not expected cell state")
	}
}

func countLiveNeighboors(board [][]int, i, j int) int {
	count := 0

	for l := max(0, i-1); l < min(len(board), i+2); l++ {
		for k := max(0, j-1); k < min(len(board[0]), j+2); k++ {
			if l == i && k == j {
				continue
			}
			if board[l][k] == live || board[l][k] == dying {
				count++
			}
		}
	}

	return count
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
