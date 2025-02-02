package countthenumberofconsistentstrings

func countConsistentStrings(allowed string, words []string) int {
	m := make(map[byte]struct{}, len(allowed))

	for i := 0; i < len(allowed); i++ {
		m[allowed[i]] = struct{}{}
	}

	count := 0

	for i := 0; i < len(words); i++ {
		word := words[i]
		isConsistent := true

		for j := 0; j < len(word); j++ {
			if _, ok := m[word[j]]; !ok {
				isConsistent = false
				break
			}
		}

		if isConsistent {
			count++
		}
	}

	return count
}
