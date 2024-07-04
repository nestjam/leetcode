package longestpalindrome

func longestPalindrome(s string) int {
	const (
		alphabet = 26 + 6 + 26
		a        = 65
	)
	m := make([]int, alphabet)

	for i := 0; i < len(s); i++ {
		m[s[i]-a]++
	}

	l := 0
	hasOdd := false
	for i := 0; i < len(m); i++ {
		if m[i]%2 == 0 {
			l += m[i]
		} else {
			if !hasOdd {
				hasOdd = true
				l++
			}
			l += m[i] - 1
		}
	}

	return l
}
