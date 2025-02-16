package deletecolumnstomakesorted

func minDeletionSize(rows []string) int {
	unsortedColumnsCount := 0

	for i := 0; i < len(rows[0]); i++ {
		if !checkIsSorted(rows, i) {
			unsortedColumnsCount++
		}
	}

	return unsortedColumnsCount
}

func checkIsSorted(rows []string, column int) bool {
	isSorted := true
	c := rows[0][column]

	for i := 1; i < len(rows); i++ {
		if c > rows[i][column] {
			isSorted = false
			break
		}
		c = rows[i][column]
	}

	return isSorted
}
