package flipcolumnsformaximumnumberofequalrows

func maxEqualRowsAfterFlips(matrix [][]int) int {
	maxCount := 0
	flips := make(map[string]int, len(matrix[0]))

	for i := 0; i < len(matrix); i++ {
		k := getKey(matrix[i])
		flips[k]++

		if flips[k] > maxCount {
			maxCount = flips[k]
		}
	}

	return maxCount
}

func getKey(r []int) string {
	b := make([]byte, len(r))

	for i := 0; i < len(r); i++ {
		if r[i] == r[0] {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}

	return string(b)
}
