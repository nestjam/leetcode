package shufflestring

func restoreString(s string, indices []int) string {
	b := make([]byte, len(s))

	for i := 0; i < len(b); i++ {
		b[indices[i]] = s[i]
	}

	return string(b)
}
