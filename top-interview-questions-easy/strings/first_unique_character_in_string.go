package strings

func firstUniqChar(s string) int {
	const a = 97
	counters := make([]int, 27)
	for i := 0; i < len(s); i++ {
		counters[s[i]-a]++
	}
	index := -1
	for i := 0; i < len(s); i++ {
		if counters[s[i]-a] == 1 {
			return i
		}
	}
	return index
}
