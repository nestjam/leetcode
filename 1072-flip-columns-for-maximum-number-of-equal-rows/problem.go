package flipcolumnsformaximumnumberofequalrows

func maxEqualRowsAfterFlips(matrix [][]int) int {
	maxCount := 0
	flipColumns := make([]bool, len(matrix[0]))
	flippedColumns := make(map[string]struct{}, len(matrix[0]))

	for i := 0; i < len(matrix); i++ {
		r := matrix[i]

		flipColumns[0] = true
		for j := 1; j < len(r); j++ {
			flipColumns[j] = r[0] == r[j]
		}

		key := getKey(flipColumns)

		if _, ok := flippedColumns[key]; ok {
			continue
		}

		flippedColumns[key] = struct{}{}

		count := countEqualRows(matrix, flipColumns)
		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}

func getKey(columns []bool) string {
	b := make([]byte, len(columns))

	for i := 0; i < len(columns); i++ {
		if columns[i] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}

	return string(b)
}

func countEqualRows(matrix [][]int, flippedColumns []bool) int {
	var count = 0

	for i := 0; i < len(matrix); i++ {
		if equalValues(matrix[i], flippedColumns) {
			count++
		}
	}

	return count
}

func equalValues(row []int, flippedColumns []bool) bool {
	value := getValue(row[0], flippedColumns[0])

	for i := 1; i < len(row); i++ {
		if value != getValue(row[i], flippedColumns[i]) {
			return false
		}
	}

	return true
}

func getValue(value int, isFlipped bool) int {
	if isFlipped {
		if value == 0 {
			return 1
		} else {
			return 0
		}
	}

	return value
}
