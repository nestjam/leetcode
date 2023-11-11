package other

func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}

	t := make([][]int, numRows)
	t[0] = []int{1}
	t[1] = []int{1, 1}

	for n := 2; n < numRows; n++ {
		next := make([]int, n+1)
		next[0] = 1
		next[len(next)-1] = 1
		for i := 1; i < len(next)-1; i++ {
			next[i] = t[n-1][i] + t[n-1][i-1]
		}
		t[n] = next
	}
	return t
}

// with recursion
/*
func generate(numRows int) [][]int {
    if numRows > 2 {
		t := generate(numRows - 1)
		prev := t[len(t) - 1]
		next := make([]int, numRows)
		next[0] = 1
		next[len(next) - 1] = 1
		for i := 1; i < numRows - 1; i++ {
			next[i] = prev[i] + prev[i - 1]
		}
		t = append(t, next)
		return t
	} else if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}
	return [][]int{{1}}
}
*/
