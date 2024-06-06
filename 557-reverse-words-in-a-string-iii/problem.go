package reversewordsinastringiii

func reverseWords(s string) string {
	b := []byte(s)

	for i, j := 0, 0; j < len(b); {
		for ; j < len(b) && b[j] != ' '; j++ {
		}
		j--

		for k, l := i, j; k < l; k, l = k+1, l-1 {
			b[k], b[l] = b[l], b[k]
		}

		j += 2
		i = j
	}

	return string(b)
}
