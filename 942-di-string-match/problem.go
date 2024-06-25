package distringmatch

func diStringMatch(s string) []int {
	n := make([]int, len(s)+1)
	l, h := 0, len(s)

	for i := 0; i < len(s); i++ {
		if s[i] == 'I' {
			n[i] = l
			l++
		} else {
			n[i] = h
			h--
		}
	}

	n[len(s)] = l

	return n
}
