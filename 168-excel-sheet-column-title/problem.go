package excelsheetcolumntitle

func convertToTitle(n int) string {
	const (
		alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		s        = 26
	)

	v := make([]byte, 0, 10)

	for {
		n--
		v = append(v, alphabet[n%s])
		n = n / s
		if n == 0 {
			break
		}
	}

	for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
		v[i], v[j] = v[j], v[i]
	}

	return string(v)
}
