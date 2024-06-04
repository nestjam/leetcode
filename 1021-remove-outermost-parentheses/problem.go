package removeoutermostparentheses

func removeOuterParentheses(s string) string {
	r := make([]byte, 0, len(s))
	sum, i := 0, 0

	for j := 0; j < len(s); j++ {
		v := 1
		if s[j] == ')' {
			v = -1
		}

		if sum == 0 {
			i = j
		}

		sum += v

		if sum == 0 {
			r = append(r, s[i+1:j]...)
		}
	}

	return string(r)
}
