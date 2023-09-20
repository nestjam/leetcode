package strings

func isPalindrome(s string) bool {
	i, j := 0, len(s) - 1
	for i <= j {
		if !isAlphanumeric(s[i]) {
			i++
			continue
		}
		if !isAlphanumeric(s[j]) {
			j--
			continue
		}
		if toLower(s[i]) != toLower(s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func toLower(b byte) byte {
	if b >= 65 && b <= 90 { //A..Z
		return b + 32 // a..z
	}
	return b
}

func isAlphanumeric(b byte) bool {
	if b >= 48 && b <= 57 { //0..9
		return true
	}
	if b >= 65 && b <= 90 { //A..Z
		return true
	}
	if b >= 97 && b <= 122 { //a..z
		return true
	}
	return false
}