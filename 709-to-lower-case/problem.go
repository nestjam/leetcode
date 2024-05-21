package tolowercase

func toLowerCase(s string) string {
	const (
		d byte = 32
		A byte = 65
		Z byte = 90
	)
	r := make([]byte, len(s))

	for i := 0; i < len(s); i++ {
		r[i] = s[i]

		if r[i] < A || r[i] > Z {
			continue
		}

		r[i] += d
	}

	return string(r)
}
