package maximumnestingdepthoftheparentheses

func maxDepth(s string) int {
	n, m := 0, 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			n++
		} else if s[i] == ')' {
			if n > m {
				m = n
			}
			n--
		}
	}

	return m
}
