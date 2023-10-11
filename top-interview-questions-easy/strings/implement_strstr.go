package strings

func strStr(haystack string, needle string) int {
	j := 0
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[j] {
			j++
			if j == len(needle) {
				return i - j + 1
			}
		} else {
			i = i - j
			j = 0
		}

	}
	return -1
}

// func strStr(haystack string, needle string) int {
// 	for i := 0; i <= len(haystack) - len(needle); i++ {
// 		j := 0
// 		for ; j < len(needle); j++ {
// 			if haystack[i + j] != needle[j] {
// 				break
// 			}
// 		}
// 		if j == len(needle) {
// 			return i
// 		}
// 	}
// 	return -1
// }