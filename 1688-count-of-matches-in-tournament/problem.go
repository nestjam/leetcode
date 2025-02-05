package countofmatchesintournament

func numberOfMatches(n int) int {
	m := 0

	for n > 1 {
		i := n / 2
		m += i
		n = i + n%2
	}

	return m
}
