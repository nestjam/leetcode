package countthenumberofconsistentstrings

func countConsistentStrings(allowed string, words []string) int {
	a := getBinarySignature(allowed)
	count := 0

	for i := 0; i < len(words); i++ {
		b := getBinarySignature(words[i])

		if a&b == b {
			count++
		}
	}

	return count
}

func getBinarySignature(word string) int {
	const a byte = 97
	mask := [26]bool{}

	for i := 0; i < len(word); i++ {
		mask[word[i]-a] = true
	}

	b := 0

	for i := 0; i < len(mask); i++ {
		if mask[i] {
			b |= 1
		}
		b = b << 1
	}

	return b
}
