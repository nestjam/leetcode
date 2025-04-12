package longestpalindromicsubstring

func longestPalindrome(s string) string {
	for l := len(s); l > 0; l-- {
		for i := 0; i <= len(s)-l; i++ {
			if isPalindrome(s[i:l+i]) {
				return s[i:l+i]
			}
		}
	}
	
	return ""
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		if s[l] != s[r] {
			return false
		}

		l++
		r--
	}

	return true
}