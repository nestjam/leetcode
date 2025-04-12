package longestpalindromicsubstring

func longestPalindrome(s string) string {
	n, m := 0, 0

	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if s[i] != s[j] {
				continue
			}

			if m-n < j-i && isPalindrome(s[i:j+1]) {
				n, m = i, j
			}
		}
	}
	
	return s[n:m+1]
}

func isPalindrome(s string) bool {
	n := len(s)

	for i := 0; i < n/2; i++ {
		if s[i] != s[n-i-1] {
			return false
		}
	}

	return true
}