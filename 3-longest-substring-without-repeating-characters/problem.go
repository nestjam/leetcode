package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
	i, j, n := 0, 0, 0
	chars := make(map[byte]int)

	for k := 0; k < len(s); k++ {
		if l, ok := chars[s[k]]; ok {
			if len(chars) > n {
				n = len(chars)
			}

			for p := i; p < l; p++ {
				delete(chars, s[p])
			}
			chars[s[l]] = k

			i, j = l+1, k
		} else {
			j++
			chars[s[k]] = k
		}
	}

	if len(chars) > n {
		n = len(chars)
	}

	return n
}
