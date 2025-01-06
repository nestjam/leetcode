package shortestdistancetoacharacter

func shortestToChar(s string, c byte) []int {
	d := make([]int, len(s))
	l := -1

	for i := 0; i < len(s); i++ {
		if s[i] == c {
			if l == -1 {
				for j := i - 1; j >= 0; j-- {
					d[j] = i - j
				}
			} else {
				for j := i - 1; j > l; j-- {
					if d[j] <= i-j {
						break
					}
					d[j] = i - j
				}
			}

			l = i
		} else if l != -1 {
			d[i] = i - l
		}
	}

	return d
}
