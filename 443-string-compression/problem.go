package stringcompression

import "strconv"

func compress(s []byte) int {
	j, n := 0, 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			n++
		} else {
			s[j] = s[i-1]
			if n > 1 {
				a := strconv.Itoa(n)
				for l := 0; l < len(a); l++ {
					j++
					s[j] = a[l]
				}
				n = 1
			}
			j++
		}
	}

	s[j] = s[len(s)-1]
	if n > 1 {
		a := strconv.Itoa(n)
		for l := 0; l < len(a); l++ {
			j++
			s[j] = a[l]
		}
	}

	return j + 1
}
