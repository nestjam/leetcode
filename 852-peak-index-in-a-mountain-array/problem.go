package peakindexinamountainarray

func peakIndexInMountainArray(arr []int) int {
	if len(arr) == 3 {
		return 1
	}

	s := make([]bool, len(arr)-1)
	for i := 0; i < len(s); i++ {
		if arr[i+1] > arr[i] {
			s[i] = true
		}
	}

	i, j := 0, len(s)-1

	for i < j-1 {
		m := (j + i) / 2

		if s[i] != s[m] {
			j = m
		} else {
			i = m
		}
	}

	return j
}
